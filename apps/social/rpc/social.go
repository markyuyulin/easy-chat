package main

import (
	"flag"
	"fmt"
	"imooc/easy-chat/pkg/interceptc/rpcserver"

	"imooc/easy-chat/apps/social/rpc/internal/config"
	"imooc/easy-chat/apps/social/rpc/internal/server"
	"imooc/easy-chat/apps/social/rpc/internal/svc"
	"imooc/easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/dev/social.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		social.RegisterSocialServer(grpcServer, server.NewSocialServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	//加入错误的拦截器
	s.AddUnaryInterceptors(rpcserver.LogInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
