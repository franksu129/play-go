package main

import (
	"GoConcurrency/changeValue"
	"GoConcurrency/waitWay"
	"fmt"
)

func main() {
	fmt.Println("--------- Run sync wait test ---------")
	testWait()
	fmt.Println("--------- End ---------")
	fmt.Println("--------- Run change value test ---------")
	testChangeValue()
	fmt.Println("--------- End ---------")
}

// 非同步等待任務的作法
func testWait() {
	// #1. sync.WaitGroup 方式
	waitWay.UseWaitGroup()
	// #2. Channel 方式
	waitWay.UseChannel()
}

// 非同步修改變數
func testChangeValue() {
	// #1. UseMutex 方式
	changeValue.UseMutex()
	// #2. Channel 方式
	changeValue.UseChannel()
}
