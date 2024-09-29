package logic

import (
	"context"
	"zero-mall/service/product/model"

	"zero-mall/service/product/rpc/internal/svc"
	"zero-mall/service/product/rpc/types/product"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	newProduct := &model.Product{}
	err := copier.Copy(newProduct, in)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	res, err := l.svcCtx.ProductModel.Insert(l.ctx, newProduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newProduct.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.CreateResponse{
		Id: newProduct.Id,
	}, nil
}
