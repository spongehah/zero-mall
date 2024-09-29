package logic

import (
	"context"
	"zero-mall/service/product/rpc/types/product"

	"zero-mall/service/product/api/internal/svc"
	"zero-mall/service/product/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	rpcReq := &product.DetailRequest{}
	copier.Copy(rpcReq, req)
	rpcResp, err := l.svcCtx.ProductRpc.Detail(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	resp = &types.DetailResponse{}
	copier.Copy(resp, rpcResp)
	return resp, nil
}
