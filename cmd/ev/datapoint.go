package main

import (
	"fmt"
	"strings"
	"time"
)

type datapoint struct {
	PID             int
	Time            time.Time
	Goroutines      int
	UserCPU         time.Duration
	UserCPUFrac     float64
	SysCPU          time.Duration
	SysCPUFrac      float64
	AllocBytes      int64
	NumObjects      int64
	TotalAllocBytes int64
	AllocationRate  float64 // bytes/s
	SysBytes        int64
	HeapSysBytes    int64
	StackSysBytes   int64
	NextGCBytes     int64
	TotalGCPause    time.Duration
	GCPause         time.Duration
	NumGC           int
}

func parseDataPoint(line string) (datapoint, bool) {
	if !strings.HasPrefix(line, "ev ") {
		return datapoint{}, false
	}

	var t int64
	var dp datapoint

	// ev pid 11866 @1459409325892: 5 gr, 1 ms user, 3 ms sys, 299 KiB alloc, 4472 objs, 299 KiB totalloc, 1636 KiB sys, 736 KiB heap, 288 KiB stack, 4096 KiB nextgc, 0 ms gcpause, 0 gcs
	_, err := fmt.Sscanf(line, "ev pid %d @%d: %d gr, %d ms user, %d ms sys, %d KiB alloc, %d objs, %d KiB totalloc, %d KiB sys, %d KiB heap, %d KiB stack, %d KiB nextgc, %d ms gcpause, %d gcs\n",
		&dp.PID, &t, &dp.Goroutines, &dp.UserCPU, &dp.SysCPU, &dp.AllocBytes, &dp.NumObjects, &dp.TotalAllocBytes, &dp.SysBytes,
		&dp.HeapSysBytes, &dp.StackSysBytes, &dp.NextGCBytes, &dp.TotalGCPause, &dp.NumGC)

	if err != nil {
		return datapoint{}, false
	}

	// Fix up units
	dp.Time = time.Unix(0, t*int64(time.Millisecond))
	dp.UserCPU *= time.Millisecond
	dp.SysCPU *= time.Millisecond
	dp.AllocBytes *= 1024
	dp.TotalAllocBytes *= 1024
	dp.SysBytes *= 1024
	dp.HeapSysBytes *= 1024
	dp.StackSysBytes *= 1024
	dp.NextGCBytes *= 1024
	dp.TotalGCPause *= time.Millisecond

	return dp, true
}
