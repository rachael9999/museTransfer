package logic

import (
	"context"
	"errors"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateResponse, err error) {
	// check if the edited name contains in folder
	cnt, err := l.svcCtx.Engine.Where("filename = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?) AND user_identity = ?", req.Filename, req.Identity, userIdentity).Count(&models.UserRepository{})
	if (cnt > 0) {
		return nil, errors.New("the name has been used")
	}

	if err != nil {
		return nil, err
	}

	// name update
	data := &models.UserRepository{
		Filename: req.Filename,
		UpdatedAt: time.Now(),
	}
	_, err = l.svcCtx.Engine.Table(&models.UserRepository{}).
	Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
	Update(data)

	if err != nil {
		return nil, err
	}
	return
}
