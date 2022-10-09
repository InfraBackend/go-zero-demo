package svc

import (
	"zero-demo/my_book_sys/user/api/internal/config"
	"zero-demo/my_book_sys/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserHttpModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserHttpModel: model.NewUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
