package logic

import (
	"context"
	"zero-mall/service/product/rpc/types/product"

	"zero-mall/service/product/api/internal/svc"
	"zero-mall/service/product/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	rpcReq := &product.CreateRequest{}
	err = copier.Copy(rpcReq, req)
	if err != nil {
		return nil, err
	}
	res, err := l.svcCtx.ProductRpc.Create(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	resp = &types.CreateResponse{}
	err = copier.Copy(resp, res)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
