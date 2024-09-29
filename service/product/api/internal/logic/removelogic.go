package logic

import (
	"context"
	"zero-mall/service/product/rpc/types/product"

	"zero-mall/service/product/api/internal/svc"
	"zero-mall/service/product/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	rpcReq := &product.RemoveRequest{}
	copier.Copy(rpcReq, req)
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return &types.RemoveResponse{}, nil
}
