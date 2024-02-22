package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)

	data := make([]*models.User, 0)
	err = l.svcCtx.Engine.Find(&data)
	if err != nil {
		log.Println("Get user error", err)
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshal user error", err)
	}

	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "    ")
	if err != nil {	
		log.Println("Indent user error", err)
	}

	fmt.Println(dst.String())

	resp.Message = dst.String()

	return
}
