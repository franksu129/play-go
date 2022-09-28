package waitWay

import (
	"fmt"
	"sync"
	"time"
)

func UseWaitGroup() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go sayTakeWG("Hello wg", wg)

	wg.Wait()
}

func sayTakeWG(s string, wg *sync.WaitGroup) {
	// defer 定義 func 結束後執行
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}
