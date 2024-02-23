package logic

import (
	"context"
	"time"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {

	rp := models.RepositoryPool{
		Identity: helper.UUID(),
		Hash: 	 	req.Hash,
		Filename: req.Filename,
		Ext: 	 		req.Ext,
		Size: 	 	req.Size,
		Path: 	 	req.Path,
		CreatedAt: time.Now(),
	}

	_, err = l.svcCtx.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}

	resp = new(types.FileUploadResponse)
	resp.Identity = rp.Identity
	resp.Filename = rp.Filename
	resp.Ext = rp.Ext

	return
}
