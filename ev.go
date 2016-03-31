package ev

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func init() {
	intv, _ := strconv.ParseInt(os.Getenv("EV_INTERVAL"), 10, 64)
	if intv > 0 {
		initTrace()
		go printTrace(time.Duration(intv) * time.Millisecond)
	}
}

func printTrace(sleep time.Duration) {
	var mem runtime.MemStats
	pid := os.Getpid()
	for {
		user, sys := cpuUsage()
		runtime.ReadMemStats(&mem)
		fmt.Printf("ev pid %d @%d: %d gr, %d ms user, %d ms sys, %d KiB alloc, %d objs, %d KiB totalloc, %d KiB sys, %d KiB heap, %d KiB stack, %d KiB nextgc, %d ms gcpause, %d gcs\n",
			pid, time.Now().UnixNano()/1e6, runtime.NumGoroutine(), user/1e6, sys/1e6, mem.Alloc/1024, mem.HeapObjects, mem.TotalAlloc/1024, (mem.Sys-mem.HeapReleased)/1024,
			(mem.HeapSys-mem.HeapReleased)/1024, mem.StackSys/1024, mem.NextGC/1024, mem.PauseTotalNs/1e6, mem.NumGC)
		time.Sleep(sleep)
	}
}
