package ev

import (
	"fmt"
	"os"
	"reflect"
	"time"
	"unsafe"
)

type timestruct struct {
	sec  int64
	nsec int64
}

func (tv timestruct) duration() time.Duration {
	return time.Duration(tv.sec*1e9 + tv.nsec)
}

type prusage struct {
	_     [4*16 + 2*4]byte
	utime timestruct
	stime timestruct
}

var fd *os.File
var rusage prusage
var buf []byte

func initTrace() {
	var pid = os.Getpid()
	fd, _ = os.Open(fmt.Sprintf("/proc/%d/usage", pid))

	*(*reflect.SliceHeader)(unsafe.Pointer(&buf)) = reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&rusage)),
		Len:  4*16 + 2*4 + 2*2*8,
		Cap:  4*16 + 2*4 + 2*2*8,
	}
}

func cpuUsage() (user, sys time.Duration) {
	_, err := fd.Seek(0, 0)
	if err != nil {
		return 0, 0
	}
	_, err = fd.Read(buf)
	if err != nil {
		return 0, 0
	}
	return rusage.utime.duration(), rusage.stime.duration()
}
