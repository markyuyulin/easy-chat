// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	friend "imooc/easy-chat/apps/social/api/internal/handler/friend"
	group "imooc/easy-chat/apps/social/api/internal/handler/group"
	"imooc/easy-chat/apps/social/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 好友申请
				Method:  http.MethodPost,
				Path:    "/friend/putIn",
				Handler: friend.FriendPutInHandler(serverCtx),
			},
			{
				// 好友申请处理
				Method:  http.MethodPut,
				Path:    "/friend/putIn",
				Handler: friend.FriendPutInHandleHandler(serverCtx),
			},
			{
				// 好友申请列表
				Method:  http.MethodGet,
				Path:    "/friend/putIns",
				Handler: friend.FriendPutInListHandler(serverCtx),
			},
			{
				// 好友列表
				Method:  http.MethodGet,
				Path:    "/friends",
				Handler: friend.FriendListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/social"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创群
				Method:  http.MethodPost,
				Path:    "/group",
				Handler: group.CreateGroupHandler(serverCtx),
			},
			{
				// 申请进群
				Method:  http.MethodPost,
				Path:    "/group/putIn",
				Handler: group.GroupPutInHandler(serverCtx),
			},
			{
				// 申请进群处理
				Method:  http.MethodPut,
				Path:    "/group/putIn",
				Handler: group.GroupPutInHandleHandler(serverCtx),
			},
			{
				// 申请进群列表
				Method:  http.MethodGet,
				Path:    "/group/putIns",
				Handler: group.GroupPutInListHandler(serverCtx),
			},
			{
				// 成员列表列表
				Method:  http.MethodGet,
				Path:    "/group/users",
				Handler: group.GroupUserListHandler(serverCtx),
			},
			{
				// 用户申群列表
				Method:  http.MethodGet,
				Path:    "/groups",
				Handler: group.GroupListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/social"),
	)
}
