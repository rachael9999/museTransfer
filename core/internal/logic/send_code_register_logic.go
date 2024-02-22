package logic

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
	bolt "go.etcd.io/bbolt"
)

type SendCodeRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}


func NewSendCodeRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeRegisterLogic {
	return &SendCodeRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeRegisterLogic) SendCodeRegister(req *types.SendCodeRequest) (resp *types.SendCodeResponse, err error) {
	user := &models.User{Email: req.Email}
	has, err := l.svcCtx.Engine.Get(user)
	if err != nil {
		return nil, err
	}
	if has {
		return &types.SendCodeResponse{Error: "The email has been registered"}, nil
	}
	
	code := helper.Code()

	codeStruct := &models.Code{
		Value:      code,
		Expiration: time.Now().Add(10 * time.Minute), // The code will expire in 10 minutes
	}

	codeJson, err := json.Marshal(codeStruct)

	if err != nil {
		return nil, err
	}

	err = l.svcCtx.BoltDB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("codes"))
		if err != nil {
			return err
		}

		err = b.Put([]byte(req.Email), []byte(codeJson))
		if err != nil {
			return err
		}

		return nil
	})


	// use a goroutine to send the code so it doesnt block the main thread
	go func(){
		err = helper.MailCode(req.Email, code)

		if err != nil {
			log.Printf("MailCode error: %v", err)
		}
	}()

	return &types.SendCodeResponse{}, nil
}
