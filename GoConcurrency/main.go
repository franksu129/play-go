package main

import (
	"GoConcurrency/waitWay"
)

// 可以在 主執行緒 結束後
// 等待任務的作法

func main() {

	// #1. sync.WaitGroup 方式
	waitWay.UseWaitGroup()
	// #2. Channel 方式
	waitWay.UseChannel()
}
