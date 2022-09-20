package changeValue

import (
	"fmt"
	"time"
)

func UseChannel() {
	total := 0
	// 宣告容量為1, 這樣channel再放入1個物件不會進入等待狀態
	ch := make(chan int, 1)
	ch <- total
	for i := 0; i < 1000; i++ {
		go func() {
			ch <- <-ch + 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
