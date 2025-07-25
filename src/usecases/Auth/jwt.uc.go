package usecases_auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUsecase interface {
	GenerateToken(userId uint64, email string) (string, error)
	ValidateToken(token string) (uint64, error)
}

type jwtUsecase struct {
	secretKey string
	issuer    string
}

func NewJWTUsecase(secretKey string, issuer string) JWTUsecase {
	return &jwtUsecase{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

func (u *jwtUsecase) GenerateToken(userId uint64, email string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"iss":   u.issuer,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
		"nbf":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(u.secretKey))
}

func (u *jwtUsecase) ValidateToken(token string) (uint64, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(u.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	return uint64(claims["sub"].(float64)), nil
}
