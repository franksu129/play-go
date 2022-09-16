package changeValue

/*
　與C# Lock相同
　使用互斥鎖的方式來控制
*/

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	v   int
	mux sync.Mutex
}

func UseMutex() {
	total := SafeNumber{v: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			total.mux.Lock()
			total.v++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	total.mux.Lock()
	fmt.Println(total.v)
	total.mux.Unlock()
}
