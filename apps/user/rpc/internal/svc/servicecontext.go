package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"imooc/easy-chat/apps/user/models"
	"imooc/easy-chat/apps/user/rpc/internal/config"
	"imooc/easy-chat/pkg/constants"
	"imooc/easy-chat/pkg/ctxdata"
	"time"
)

type ServiceContext struct {
	Config config.Config
	*redis.Redis
	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		Redis:      redis.MustNewRedis(c.Redisx),
		UsersModel: models.NewUsersModel(sqlConn, c.Cache),
	}
}

func (svc *ServiceContext) SetRootToken() error {
	// 生成jwt
	systemToken, err := ctxdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(), 999999999, constants.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}
	// 写入到redis,k:v形式，将token存入到redis当中

	return svc.Redis.Set(constants.REDIS_SYSTEM_ROOT_TOKEN, systemToken)
}
