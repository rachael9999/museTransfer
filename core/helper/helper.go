package helper

import (
	"crypto/md5"
	"fmt"

	"cloud-disk/core/define"

	"crypto/tls"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
	"github.com/golang-jwt/jwt/v4"
	
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