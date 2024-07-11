package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
	secretKey string
}

func NewJwtService() *jwtService {
	secretKey := os.Getenv("SECRETKEYJWT")
	if secretKey == "" {
		log.Fatalf("JWT_SECRET is not set in the environment variables")
	}

	return &jwtService{
		secretKey: secretKey,
	}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
