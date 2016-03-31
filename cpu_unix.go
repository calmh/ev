//+build !windows,!solaris

package ev

import (
	"syscall"
	"time"
)

func initTrace() {}

var rusage syscall.Rusage

func cpuUsage() (user, sys time.Duration) {
	syscall.Getrusage(syscall.RUSAGE_SELF, &rusage)
	return time.Duration(rusage.Utime.Nano()), time.Duration(rusage.Stime.Nano())
}
