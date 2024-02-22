package helper

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"path"

	"cloud-disk/core/define"

	"crypto/tls"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name: 	 name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	codeStr := fmt.Sprintf("%d", code)
	return codeStr
}

func MailCode(receiverEmail, code string) error {
	e := email.NewEmail()
	e.From = "test <13816644982@163.com>"
	e.To = []string{receiverEmail}
	e.Bcc = []string{"test_bcc@example.com"}

	e.Subject = "Your code is: " + code
	e.HTML = []byte("Your code is: <h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "13816644982@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

	if err != nil {
		return err
	}
	return nil

}

func UUID() string {
	return uuid.NewV4().String()
}

func UploadFile(r *http.Request) (string, error) {
	log.Println("Cos bucket" + define.CosBucket)
	u, _ := url.Parse(define.CosBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
					SecretID:  define.CosSecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
					SecretKey: define.CosSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			},
	})

	file, fileHeader, err := r.FormFile("file")
	name := "cloud-disk/" + UUID() + path.Ext(fileHeader.Filename)
	_, err = c.Object.Put(context.Background(), name, file, nil)
	if err != nil {
			panic(err)
	}
	return define.CosBucket + "/" + name, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {

	uc := new (define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, errors.New("token is invalid")
	}
	return uc, err

}