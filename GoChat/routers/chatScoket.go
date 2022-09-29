package routers

import (
	serviceChat "GoChat/service"

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

	m.HandleConnect(func(s *melody.Session) {
		name := s.Request.URL.Query().Get("id")
		serviceChat.JoinIn(name, s)

		bMsg := serviceChat.GetBroadcastMsg(name, " is join in us")
		m.BroadcastOthers(bMsg, s)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		userInfo := serviceChat.GetUserInfo(s)
		bMsg := serviceChat.GetBroadcastMsg(userInfo.Name, "leave us")
		m.BroadcastOthers(bMsg, s)

	})
}
