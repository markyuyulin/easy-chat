package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"imooc/easy-chat/apps/social/api/internal/config"
	"imooc/easy-chat/apps/social/rpc/socialclient"
	"imooc/easy-chat/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	//社交RPC的客户端
	Social socialclient.Social
	User   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
