package logic

import (
	"context"
	"time"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepoSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepoSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepoSaveLogic {
	return &UserRepoSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepoSaveLogic) UserRepoSave(req *types.UserRepoSaveRequest, Identity string) (resp *types.UserRepoSaveResponse, err error) {
	ur := &models.UserRepository{	
		Identity: 							helper.UUID(),
		UserIdentity: 					Identity,
		ParentId: 							req.ParentId,
		RepositoryIdentity: 		req.RepositoryIdentity,
		Ext: 										req.Ext,
		Filename: 							req.Filename,
		CreatedAt: 							time.Now(),
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return
	}
	return &types.UserRepoSaveResponse{
		Identity: ur.Identity,
	}, nil

}
