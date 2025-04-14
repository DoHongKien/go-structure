package auth

import (
	"time"

	"github.com/DoHongKien/go-structure/global"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

func CreateAccessToken(uuidToken string) (string, error) {
	expireTime := global.Config.JWT.ExpireTime
	return CreateToken(uuidToken, expireTime)
}

func CreateRefreshToken(uuidToken string) (string, error) {
	expireTime := global.Config.JWT.RefreshTime
	return CreateToken(uuidToken, expireTime)
}

func CreateToken(uuidToken string, expireTime string) (string, error) {

	if expireTime == "" {
		expireTime = "1h"
	}

	expiration, err := time.ParseDuration(expireTime)

	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(expiration)

	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id: 	  uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt: now.Unix(),
			Issuer: "kiendh",
			Subject: uuidToken,
		},
	})
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return claims.SignedString([]byte(global.Config.JWT.SecretKey))
}

func VerifyTokenSubject(token string) (*jwt.StandardClaims, error) {
	claims, err := ParseJwtTokenSubject(token)

	if err != nil {
		return nil, err
	}

	if err = claims.Valid(); err != nil {
		return nil, err
	}

	return claims, nil
}

func ParseJwtTokenSubject(token string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.SecretKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
