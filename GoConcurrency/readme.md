# GoConcurrency
主要測試並行時緒之間的等待和處理

### changeValue
修改共通參數時的作法
1. Channel
2. sync.Mutex (互斥鎖)

### waitWay
等待任務結束作法
1. Channel
2. sync.WaitGroup (任務數量控制)