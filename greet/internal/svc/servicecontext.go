package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/greet/internal/config"
	"go-zero-demo/greet/internal/logic/model"

)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
	AccountModel model.AccountModel
	DebtModel model.DebtModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		AccountModel: model.NewAccountModel(conn, c.CacheRedis),

	}
}
