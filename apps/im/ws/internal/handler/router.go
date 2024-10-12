package handler

import (
	"imooc/easy-chat/apps/im/ws/internal/handler/conversation"
	"imooc/easy-chat/apps/im/ws/internal/handler/user"
	"imooc/easy-chat/apps/im/ws/internal/svc"
	"imooc/easy-chat/apps/im/ws/websocket"
)

// 路由的加载
func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
	})
}
