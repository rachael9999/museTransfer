package logic

import (
	"context"

	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthLogic {
	return &RefreshAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthLogic) RefreshAuth(req *types.RefreshAuthRequest) (resp *types.RefreshAuthResponse, err error) {
	claims, err := helper.ParseToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	// regenerate new token
	token, err := helper.GenerateToken(claims.Id, claims.Identity, claims.Name, define.TokenExpireTime)
	if err != nil {
		return nil, err
	}
	
	RefreshToken, err := helper.GenerateToken(claims.Id, claims.Identity, claims.Name, define.TokenExpireTime*2)

	if err != nil {
		return nil, err
	}

	resp = new(types.RefreshAuthResponse)
	resp.Token = token
	resp.RefreshToken = RefreshToken

	return
}
