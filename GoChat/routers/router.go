package routers

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()
	setStaticWeb(r)
	setChatSocket(r)
	SetSystemController(r)

	return r
}
