package logic

import (
	"context"
	"zero-mall/service/product/rpc/types/product"

	"zero-mall/service/product/api/internal/svc"
	"zero-mall/service/product/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	rpcReq := &product.UpdateRequest{}
	copier.Copy(rpcReq, req)
	rpcResp, err := l.svcCtx.ProductRpc.Update(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	resp = &types.UpdateResponse{}
	copier.Copy(resp, rpcResp)
	return resp, nil
}
