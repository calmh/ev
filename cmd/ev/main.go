//go:generate go run ../../script/genassets.go -o assets.go ../../gui

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	_ "github.com/calmh/ev"
)

type dataset struct {
	Time   []int64 // Unix milliseconds, for Javascript
	Series []series
}

type series struct {
	Name   string
	Unit   string
	Values []float64
}

type request struct {
	pid      int
	start    time.Time
	end      time.Time
	pids     chan<- []int
	response chan<- response
}

type response struct {
	Dataset dataset
	Lines   []logline
}

type logline struct {
	Time time.Time
	Line string
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "Listen address")
	keep := flag.Duration("keep", 5*time.Minute, "Length of history to keep")
	flag.Parse()

	reqs := make(chan request)
	go handler(reqs, *keep)

	h := &httpHandler{
		reqs: reqs,
	}

	http.HandleFunc("/", handleStatic)
	http.HandleFunc("/data", h.handleRequest)
	http.ListenAndServe(*addr, nil)
}

type httpHandler struct {
	reqs chan<- request
}

func (h *httpHandler) handleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pid, _ := strconv.ParseInt(r.Form.Get("pid"), 10, 32)
	start, _ := strconv.ParseInt(r.Form.Get("start"), 10, 64)
	end, _ := strconv.ParseInt(r.Form.Get("end"), 10, 64)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if pid > 0 {
		resp := make(chan response)
		h.reqs <- request{
			pid:      int(pid),
			start:    time.Unix(start/1000, (start%1000)*1e6),
			end:      time.Unix(end/1000, (end%1000)*1e6),
			response: resp,
		}
		res := <-resp
		json.NewEncoder(w).Encode(res)
	} else {
		pids := make(chan []int)
		h.reqs <- request{
			pids: pids,
		}
		res := <-pids
		json.NewEncoder(w).Encode(res)
	}
}

func handler(reqs <-chan request, keep time.Duration) {
	c := make(chan datapoint)
	l := make(chan logline)
	go datapointsInto(os.Stdin, c, l)

	dps := make(map[int][]datapoint)
	var lls []logline

	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case dp := <-c:
			l := dps[dp.PID]

			cut := 0
			for cut = range l {
				if time.Since(l[cut].Time) < keep {
					break
				}
			}
			l = l[cut:]

			l = append(l, dp)
			dps[dp.PID] = l

		case ll := <-l:
			lls = append(lls, ll)

		case <-t.C:
			cut := 0
			for cut = range lls {
				if time.Since(lls[cut].Time) < keep {
					break
				}
			}
			lls = lls[cut:]

		case req := <-reqs:
			if req.pid > 0 {
				points := dps[req.pid]
				if req.start.Unix() > 0 && req.end.After(req.start) {
					points = filterPoints(req.start, req.end, points)
				}
				points = reducePoints(250, points)
				ds := pointsToSeries(points)
				lines := lls
				if req.start.Unix() > 0 && req.end.After(req.start) {
					lines = filterLines(req.start, req.end, lines)
				}
				if len(lines) > 100 {
					lines = lines[:100]
				}
				req.response <- response{
					Dataset: ds,
					Lines:   lines,
				}
			} else {
				pids := make([]int, 0, len(dps))
				for pid := range dps {
					pids = append(pids, pid)
				}
				sort.Ints(pids)
				req.pids <- pids
			}
		}
	}
}

func datapointsInto(r io.Reader, c chan<- datapoint, l chan<- logline) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if dp, ok := parseDataPoint(line); ok {
			c <- dp
		} else {
			fmt.Printf("%s\n", s.Bytes())
			l <- logline{
				Time: time.Now(),
				Line: s.Text(),
			}
		}
	}
}

func filterPoints(start, end time.Time, points []datapoint) []datapoint {
	var si, ei int
	for si = 0; si < len(points); si++ {
		if points[si].Time.After(start) {
			break
		}
	}
	for ei = si; ei < len(points); ei++ {
		if points[ei].Time.After(end) {
			break
		}
	}
	return points[si:ei]
}

func filterLines(start, end time.Time, lines []logline) []logline {
	var si, ei int
	for si = 0; si < len(lines); si++ {
		if lines[si].Time.After(start) {
			break
		}
	}
	for ei = si; ei < len(lines); ei++ {
		if lines[ei].Time.After(end) {
			break
		}
	}
	return lines[si:ei]
}

func reducePoints(num int, points []datapoint) []datapoint {
	reducer := len(points)/num + 1
	result := make([]datapoint, 0, len(points)/reducer)
	for i := 0; i < len(points); i += reducer {
		result = append(result, points[i])
	}
	return result
}

func pointsToSeries(points []datapoint) dataset {
	var ds dataset
	ds.Series = []series{
		series{Name: "User CPU", Unit: "%"},
		series{Name: "System CPU", Unit: "%"},
		series{Name: "Memory in use", Unit: "B"},
		series{Name: "Memory allocated", Unit: "B"},
		series{Name: "Goroutines", Unit: ""},
		series{Name: "Allocation Rate", Unit: "bytes/s"},
		series{Name: "Next GC", Unit: "bytes"},
	}

	var prev datapoint
	for _, dp := range points {
		if prev.Time.IsZero() {
			prev = dp
			continue
		}

		td := dp.Time.Sub(prev.Time)
		dp.UserCPUFrac = float64(dp.UserCPU-prev.UserCPU) / float64(td)
		dp.SysCPUFrac = float64(dp.SysCPU-prev.SysCPU) / float64(td)
		dp.GCPause = dp.TotalGCPause - prev.TotalGCPause
		dp.AllocationRate = float64(dp.TotalAllocBytes-prev.TotalAllocBytes) / dp.Time.Sub(prev.Time).Seconds()
		ds.Time = append(ds.Time, dp.Time.UnixNano()/1e6)
		ds.Series[0].Values = append(ds.Series[0].Values, dp.UserCPUFrac*100)
		ds.Series[1].Values = append(ds.Series[1].Values, dp.SysCPUFrac*100)
		ds.Series[2].Values = append(ds.Series[2].Values, float64(dp.AllocBytes))
		ds.Series[3].Values = append(ds.Series[3].Values, float64(dp.SysBytes))
		ds.Series[4].Values = append(ds.Series[4].Values, float64(dp.Goroutines))
		ds.Series[5].Values = append(ds.Series[5].Values, dp.AllocationRate)
		ds.Series[6].Values = append(ds.Series[6].Values, float64(dp.NextGCBytes))
		prev = dp
	}
	return ds
}
