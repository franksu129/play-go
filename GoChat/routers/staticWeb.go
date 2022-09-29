package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setStaticWeb(r *gin.Engine) {
	r.LoadHTMLFiles("web/index.html")
	r.Static("/assets", "./web/assets")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
