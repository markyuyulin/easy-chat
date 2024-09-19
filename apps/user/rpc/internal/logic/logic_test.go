package logic

import (
	"github.com/zeromicro/go-zero/core/conf"
	"imooc/easy-chat/apps/user/rpc/internal/config"
	"imooc/easy-chat/apps/user/rpc/internal/svc"
)

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad("../../etc/dev/user.yaml", &c)
	svcCtx = svc.NewServiceContext(c)
}
