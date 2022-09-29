package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func setChatSocket(r *gin.Engine) {

	m := melody.New()
	r.GET("/ws", func(ctx *gin.Context) {
		m.HandleRequest(ctx.Writer, ctx.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
}
