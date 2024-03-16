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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, Identity string) (resp *types.UserFolderCreateResponse, err error) {
	// check if the edited name contains in folder
	cnt, err := l.svcCtx.Engine.Where("filename = ? AND parent_id = ? AND user_identity = ?", req.Filename, req.ParentId, Identity).Count(&models.UserRepository{})
	if (cnt > 0) {
		return nil, errors.New("the name has been used")
	}
	
	if err != nil {
		return nil, err
	}

	// crt folder
	data := &models.UserRepository{
		Identity: helper.UUID(),
		UserIdentity: Identity,
		ParentId: req.ParentId,
		Filename: req.Filename,
		CreatedAt: time.Now(),
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}

	return &types.UserFolderCreateResponse{
		Identity: data.Identity,
	}, nil
}
