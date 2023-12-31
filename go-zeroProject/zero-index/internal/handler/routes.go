// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"zore-order/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Login},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/order/info",
					Handler: orderInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/order/update",
					Handler: orderUpdateHandler(serverCtx),
				},
			}...,
		),
	)
}
