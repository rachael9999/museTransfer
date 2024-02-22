package test

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"testing"
	"time"
	"cloud-disk/core/define"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "test <13816644982@163.com>"
	e.To = []string{"rachelma4869@gmail.com"}
	e.Bcc = []string{"test_bcc@example.com"}
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	codeStr := fmt.Sprintf("%d", code)
	e.Subject = "Your code is: " + codeStr
	e.HTML = []byte("Your code is: <h1>" + codeStr + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "13816644982@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})

	if err != nil {
		t.Error(err)
	}

	}