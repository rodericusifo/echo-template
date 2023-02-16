package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/rodericusifo/echo-template/internal/pkg/config"
	"github.com/rodericusifo/echo-template/internal/pkg/types"
)

func GenerateJWTTokenFromClaims(claims *types.JwtCustomClaims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Env.JWTExpiredDuration)),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(config.Env.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.Env.PasswordHashingHashSalt)
	return string(bytes), err
}
