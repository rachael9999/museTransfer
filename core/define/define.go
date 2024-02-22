package define

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type UserClaim struct {
	Id       int
	Identity string
	Name string
	jwt.StandardClaims
}

var (
	JwtKey       string
	MailPassword string
	CosSecretId  string
	CosSecretKey string
	CosBucket    string
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}

	JwtKey = os.Getenv("JWTKey")
	MailPassword = os.Getenv("MailPassword")
	CosSecretId = os.Getenv("CosSecretId")
	CosSecretKey = os.Getenv("CosSecretKey")
	CosBucket = os.Getenv("CosBucket")
}
