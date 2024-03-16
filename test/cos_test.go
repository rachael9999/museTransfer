package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUploadByFilepath(t *testing.T) {
	u, _ := url.Parse("https://musetransfer-1324489521.cos.na-ashburn.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosSecretId,
			SecretKey: define.CosSecretKey,
		},
	})

	key := "cloud-disk/exampleobject.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./test/1.png", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse("https://musetransfer-1324489521.cos.na-ashburn.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosSecretId,
			SecretKey: define.CosSecretKey,
		},
	})

	key := "cloud-disk/exampleobject2.png"

	f, err := os.ReadFile("./test/1.png")
	if err != nil {
		return
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
}

// 分片上传初始化
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse("https://musetransfer-1324489521.cos.na-ashburn.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosSecretId,
			SecretKey: define.CosSecretKey,
		},
	})
	key := "cloud-disk/exampleobject2.png"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	t.Log(v)
	UploadID := v.UploadID // 16516751077a264ca9b5c8f560c1d1b5ea5e9d242dee047f62b02b4958491c0d90aa167d51
	fmt.Println(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse("https://musetransfer-1324489521.cos.na-ashburn.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosSecretId,
			SecretKey: define.CosSecretKey,
		},
	})
	key := "cloud-disk/exampleobject.jpeg"
	UploadID := "16516751077a264ca9b5c8f560c1d1b5ea5e9d242dee047f62b02b4958491c0d90aa167d51"
	f, err := os.ReadFile("0.chunk") // md5 : 108e92d35fe1695fbf29737d0b24561d
	if err != nil {
		t.Fatal(err)
	}
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag)
}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://musetransfer-1324489521.cos.na-ashburn.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.CosSecretId,
			SecretKey: define.CosSecretKey,
		},
	})
	key := "cloud-disk/exampleobject.jpeg"
	UploadID := "16516751077a264ca9b5c8f560c1d1b5ea5e9d242dee047f62b02b4958491c0d90aa167d51"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "108e92d35fe1695fbf29737d0b24561d"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}