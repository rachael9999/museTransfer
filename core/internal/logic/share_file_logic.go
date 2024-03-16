package logic

import (
	"context"
	"errors"
	"time"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileLogic {
	return &ShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileLogic) ShareFile(req *types.UserFileShareRequest, Identity string) (resp *types.UserFileShareResponse, err error) {
	uuid := helper.UUID()

	ur := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Table("user_repository").Where("identity = ?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return 
	}
	if !has {
		err = errors.New("user repository not found")
		return
	}
	
	data := &models.ShareBasic{
		Identity: uuid,
		UserIdentity: Identity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity: ur.RepositoryIdentity,
		ExpireTime: req.ExpireTime,
		CreatedAt: time.Now(),

	}

	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}
	resp = &types.UserFileShareResponse{
		Identity: uuid,
	}

	return
}
