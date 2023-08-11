package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zore-order/internal/config"
	"zore-order/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Login  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Login:  middleware.NewLoginMiddleware().Handle,
	}
}
