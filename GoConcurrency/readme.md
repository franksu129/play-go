# GoConcurrency
主要測試並行時緒之間的等待和處理
* changeValue: Goroutine底下修改共同參數
* waitWay: 等待Goroutine

## 修改共通參數
1. Channel
2. sync.Mutex (互斥鎖)

## 等待任務
1. Channel
2. sync.WaitGroup (任務數量控制)