package logic

import (
	"context"
	"encoding/json"
	"time"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req *types.RegisterUserRequest) (resp *types.RegisterUserResponse, err error) {
	var codeStruct models.Code

	// get the code from boltDB
	err = l.svcCtx.BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("codes"))
		codeJson := b.Get([]byte(req.Email))

		err := json.Unmarshal(codeJson, &codeStruct)
		if err != nil {
			return err
		}
		return nil
	})

	userName := &models.User{Name: req.Name}

	has, err := l.svcCtx.Engine.Get(userName)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, errors.New("the name has been registered")
	}

	user := &models.User{Email: req.Email}
	has, err = l.svcCtx.Engine.Get(user)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, errors.New("the email has been registered")
	}
	

	if err != nil {
		return nil, err
	}

	// check if the code is expired
	if time.Now().After(codeStruct.Expiration) {
		return nil, errors.New("the code is expired")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
			return nil, err
	}

	user = &models.User{
		Identity:  uuid.New().String(),
		Name:      req.Name,
		Password:  string(hashedPassword),
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	_, err = l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}


	return &types.RegisterUserResponse{}, nil
}
