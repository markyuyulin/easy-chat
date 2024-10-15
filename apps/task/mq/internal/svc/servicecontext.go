package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"imooc/easy-chat/apps/im/immodels"
	"imooc/easy-chat/apps/im/ws/websocket"
	"imooc/easy-chat/apps/task/mq/internal/config"
	"imooc/easy-chat/pkg/constants"
	"net/http"
)

type ServiceContext struct {
	config.Config

	WsClient websocket.Client
	*redis.Redis

	immodels.ChatLogModel
	//immodels.ConversationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:       c,
		Redis:        redis.MustNewRedis(c.Redisx),
		ChatLogModel: immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
		//ConversationModel: immodels.MustConversationModel(c.Mongo.Url, c.Mongo.Db),
	}

	token, err := svc.GetSystemToken()
	if err != nil {
		panic(err)
	}

	header := http.Header{}
	header.Set("Authorization", token)
	svc.WsClient = websocket.NewClient(c.Ws.Host, websocket.WithClientHeader(header))
	return svc
}

// 得到kafka的token
func (svc *ServiceContext) GetSystemToken() (string, error) {
	return svc.Redis.Get(constants.REDIS_SYSTEM_ROOT_TOKEN)
}
