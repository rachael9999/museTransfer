package logic

import (
	"context"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	user := new(models.User)
	has, err := l.svcCtx.Engine.Where("name = ?", req.Name).Get(user)
	if err != nil {
			return nil, err
	}
	if !has {
			return nil, errors.New("username not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
			return nil, errors.New("incorrect password")
	}

	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)

	if err != nil {
		return nil, err
	}

	resp = new(types.LoginReply)
	resp.Token = token
	return
}
