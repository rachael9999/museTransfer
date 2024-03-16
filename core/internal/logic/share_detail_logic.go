package logic

import (
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareDetailLogic {
	return &ShareDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareDetailLogic) ShareDetail(req *types.ShareDetailRequest) (resp *types.ShareDetailResponse, err error) {
	// update share num_count
	l.svcCtx.Engine.ShowSQL(true)
	_,err = l.svcCtx.Engine.Exec("UPDATE share SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if err != nil {
		err = errors.New("update share click_num error" + err.Error())
		return 
	}

	resp = new(types.ShareDetailResponse)
	_, err = l.svcCtx.Engine.Table("share").
		Select("user_repository.filename, repository_pool.ext, repository_pool.size, repository_pool.path, share.repository_identity").
		Join("LEFT", "repository_pool", "share.repository_identity = repository_pool.identity").
		Join("Left", "user_repository", "user_repository.identity = share.user_repository_identity").
		Where("share.identity = ?", req.Identity).Get(resp)

	if err != nil {
		return
	}
	return
}
