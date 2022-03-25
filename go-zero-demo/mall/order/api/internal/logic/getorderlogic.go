package logic

import (
	"context"
	"errors"
	"go-zero-demo/mall/user/rpc/types/user"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{Id: "1"})
	if err != nil {
		return nil, err
	}

	if user.Name != "fzdwx" {
		return nil, errors.New("user is not exist!")
	}

	return &types.OrderReply{
		Id:   req.Id,
		Name: "this is test order",
	}, nil
}
