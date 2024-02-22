package logic

import (
	"context"
	"log"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeRequest) (resp *types.SendCodeResponse, err error) {
	code := helper.Code()

	// use a goroutine to send the code so it doesnt block the main thread
	go func(){
		err = helper.MailCode(req.Email, code)

		if err != nil {
			log.Printf("MailCode error: %v", err)
		}
	}()

	return &types.SendCodeResponse{}, nil
}
