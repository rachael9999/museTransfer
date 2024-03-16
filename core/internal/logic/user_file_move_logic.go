package logic

import (
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, Identity string) (resp *types.UserFileMoveResponse, err error) {
	parentData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentId, Identity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("folder not exist")
	}

	// update parentId
	// l.svcCtx.Engine.ShowSQL(true)
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(&models.UserRepository{ParentId: int (parentData.Id)})
	if err != nil {
		return nil, err
	}

	return
}
