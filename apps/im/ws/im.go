package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"imooc/easy-chat/apps/im/ws/internal/config"
	"imooc/easy-chat/apps/im/ws/internal/handler"
	"imooc/easy-chat/apps/im/ws/internal/svc"
	"imooc/easy-chat/apps/im/ws/websocket"
	"time"
)

var configFile = flag.String("f", "etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn,
		websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)),
		// 100s心跳检测
		websocket.WithServerMaxConnectionIdle(600*time.Second),
	)
	defer srv.Stop()
	handler.RegisterHandlers(srv, ctx)

	fmt.Println("start websocket server at ", c.ListenOn, "....")
	srv.Start()
}