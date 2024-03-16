package logic

import (
	"context"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, Identity string) (resp *types.UserFileDeleteResponse, err error) {
	_, err = l.svcCtx.Engine.Table(&models.UserRepository{}).
		Where("identity = ? AND user_identity = ?", req.Identity, Identity).
		Update(map[string]interface{}{"deleted_at": time.Now()})
	
		if err != nil {
		return nil, err
	}
	return
}
