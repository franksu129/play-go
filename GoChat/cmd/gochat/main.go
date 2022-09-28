package main

import (
	"GoCaht/internal/chatInit"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

var modelConfig *chatInit.Config

func getRootPath() *string {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	return &basepath
}

func init() {
	// 因為指令執行和 VsCode的 Debug launch 位置會不同
	configPath := *getRootPath() + "/../../config"

	// 兩種方式都可以拿到Config設定，
	// viperConfig = chatInit.GetViperConfig("app", "yaml", configPath)
	// fmt.Println(viperConfig.GetInt("port"))
	modelConfig = chatInit.GetConfig("app", "yaml", configPath)
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	r.Static("/assets", "./templates/assets")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(fmt.Sprintf(":%d", modelConfig.Port))
}
