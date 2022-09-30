# 聊天室功能
主要為練習專案建立一個簡單的聊天室，並依據資料夾分類不同的功能。

主要參考來源為[random_anonymous_chat](https://github.com/codingXiang/random_anonymous_chat)，後續再針對自己的風格下去做調整。

## 相關指令
### 執行
```
go run .\cmd\gochat\
```
### image 建立/執行
```
docker build -t go-chat . 
docker run --rm -d -p 5000:5000 go-chat   
```

## 資料夾分類
* pkg        - 共享庫, 包含一切其它項目的引用代碼。
* cmd        - 主要的應用程式執行。
* config     - 設定檔。
* controller - webapi服務。
* internal   - 應用程序相關代碼(middleware、model、repository...等等)。
* routers    - 路由相關設定。
* service    - 相關服務業務邏輯。
* web        - 靜態的網站檔案。
