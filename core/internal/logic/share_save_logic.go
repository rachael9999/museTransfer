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

type ShareSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareSaveLogic {
	return &ShareSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareSaveLogic) ShareSave(req *types.UserFileSaveRequest, Identity string) (resp *types.UserFileSaveResponse, err error) {
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Table("repository_pool").Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("repository not found")
	}

	ur := &models.UserRepository{
		Identity: helper.UUID(),
		UserIdentity: Identity,
		ParentId: req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext: rp.Ext,
		Filename: rp.Filename,
		CreatedAt: time.Now(),
	}

	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	
	resp = new(types.UserFileSaveResponse)
	resp.Identity = ur.Identity

	return
}
