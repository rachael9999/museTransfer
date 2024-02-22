package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// write file into byte
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return 
		}
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// get the unique identity of the file
		hash := fmt.Sprintf("%x", md5.Sum(b))

		// check if the file is already uploaded
		rp := new(models.RepositoryPool)
		get, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)

		if err != nil {
			httpx.Error(w, err)
			return
		}

		if get {
			httpx.OkJson(w, &types.FileUploadResponse{Identity: rp.Identity})
			return 
		}

		// upload file to cos
		cosPath, err := helper.UploadFile(r)
		if err != nil {
			return
		}

		// prepare file info for upload_logic
		req.Filename = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Hash = hash
		req.Path = cosPath
		req.Size = fileHeader.Size


		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
