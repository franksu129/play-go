package routers

import (
	"GoChat/controller"

	"github.com/gin-gonic/gin"
)

func SetSystemController(r *gin.Engine) {

	r.GET("/check", controller.HealthCheck)
}
