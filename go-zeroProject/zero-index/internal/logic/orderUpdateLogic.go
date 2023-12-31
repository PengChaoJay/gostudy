package logic

import (
	"context"

	"zore-order/internal/svc"
	"zore-order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderUpdateLogic {
	return &OrderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderUpdateLogic) OrderUpdate(req *types.OrderInfoReq) (resp *types.OrderInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
