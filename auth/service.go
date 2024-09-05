package auth

import (
	"errors"
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
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(s.secretKey), nil
	})

	if err != nil {
		return tkn, err
	}

	return tkn, nil
}
