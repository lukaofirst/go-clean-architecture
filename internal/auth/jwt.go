package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	Secret         string
	AccessTokenTTL time.Duration
	Issuer         string
}

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(cfg Config, userID string) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.Issuer,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(cfg.AccessTokenTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Secret))
}

func ValidateAccessToken(cfg Config, tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		// Enforce HS256
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cfg.Secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Optional: issuer check
	if cfg.Issuer != "" && claims.Issuer != cfg.Issuer {
		return nil, errors.New("invalid issuer")
	}

	return claims, nil
}
