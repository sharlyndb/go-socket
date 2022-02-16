/**
 * @Time: 2022/2/17 01:41
 * @Author: yt.yin
 */

package limit

import (
	"log"
	"sync"
	"syscall"
)

var (

	rLimit syscall.Rlimit
	once sync.Once
)

// SetLimit 设置限流
func SetLimit() {
	once.Do(setLimit)
}

func setLimit() {
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}
