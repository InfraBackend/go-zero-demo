package svc

import (
	"zero-demo/my_book_sys/book/api/internal/config"
	"zero-demo/my_book_sys/book/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	BookHttpModel model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BookHttpModel: model.NewBookModel(sqlx.NewMysql(c.DataSource)),
	}
}
