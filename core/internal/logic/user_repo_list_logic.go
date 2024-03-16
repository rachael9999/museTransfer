package logic

import (
	"context"

	"cloud-disk/core/define"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepoListLogic {
	return &UserRepoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepoListLogic) UserRepoList(req *types.UserRepoListRequest, Identity string) (resp *types.UserRepoListResponse, err error) {
	ufl := make([]*types.UserRepoList, 0)
	resp = new(types.UserRepoListResponse)
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	pageOffset := (page - 1) * size

	l.svcCtx.Engine.ShowSQL(true)
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ? ", req.Id, Identity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity," +
			"user_repository.ext, user_repository.filename, repository_pool.path, repository_pool.size").
		Join("Left", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at IS NULL").
		Limit(size, pageOffset).Find(&ufl)

	if err != nil {
		return nil, err
	}

	count, err :=l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ? AND (deleted_at IS NULL)", req.Id, Identity).Count(&models.UserRepository{})	
	if err != nil {
		return nil, err
	}
	resp.List = ufl
	resp.Count = count

	return
}
