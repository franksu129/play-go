package main

import (
	"GoChat/internal/chatInit"
	"GoChat/routers"
	"fmt"
	"path/filepath"
	"runtime"
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
	fmt.Println("-------------------------")
	fmt.Println("$HOME")
	// 因為指令執行和 VsCode的 Debug launch 位置會不同
	configPath := *getRootPath() + "/../../config"

	// 兩種方式都可以拿到Config設定，
	// viperConfig = chatInit.GetViperConfig("app", "yaml", configPath)
	// fmt.Println(viperConfig.GetInt("port"))
	modelConfig = chatInit.GetConfig("app", "yaml", configPath)
}

func main() {
	r := routers.InitRoutes()
	r.Run(fmt.Sprintf(":%d", modelConfig.Port))
}
