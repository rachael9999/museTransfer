package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       int
	Identity string
	Name string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

var MailPassword = "EZVXMNYYDFRIZAYJ"