package test

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"cloud-disk/core/define"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestUpload(t *testing.T) {
		u, _ := url.Parse("https://musetransfer-1324489521.cos.na-ashburn.myqcloud.com")
    b := &cos.BaseURL{BucketURL: u}
    c := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  define.CosSecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
            SecretKey: define.CosSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
        },
    })

    // 对象键（Key）是对象在存储桶中的唯一标识。
    // 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
    name := "cloud-disk/test.png"
    _, _, err := c.Object.Upload(context.Background(), name, "./test/Screenshot 2024-02-17 220458.png", nil)
    if err != nil {
        panic(err)
    }
}

func TestFileUploadPut(t *testing.T) {
	u, _ := url.Parse("https://musetransfer-1324489521.cos.na-ashburn.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
					SecretID:  define.CosSecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
					SecretKey: define.CosSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			},
	})

	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "cloud-disk/test2.png"

	f, err := os.ReadFile("./test/Screenshot 2024-02-17 220458.png")
	if err != nil {
			panic(err)
	}
	_, err = c.Object.Put(context.Background(), name, bytes.NewReader(f), nil)
	if err != nil {
			panic(err)
	}
}