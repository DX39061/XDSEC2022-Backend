package controller

import (
	"XDSEC2022-Backend/src/cache"
	"XDSEC2022-Backend/src/config"
	"XDSEC2022-Backend/src/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

type TokenClaims struct {
	UserID uint   `json:"user-id"`
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}

type Token struct {
	SigningKey []byte
}

func newToken() *Token {
	return &Token{
		[]byte(config.TokenConfig.SigningKey),
	}
}

func (j *Token) createToken(claims TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *Token) parseToken(token string) (*TokenClaims, error) {
	res, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if res != nil {
		if claims, ok := res.Claims.(*TokenClaims); ok && res.Valid {
			return claims, nil
		}
		return nil, errors.New("token is not valid or claims broken")
	}
	return nil, errors.New("token is invalid")
}

//goland:noinspection GoUnusedExportedFunction
func ExtractToken(ctx *gin.Context) error {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		return errors.New("token is empty")
	}
	tokenStr := strings.ReplaceAll(token, "Bearer ", "")
	j := newToken()
	claims, err := j.parseToken(tokenStr)
	if err != nil {
		return err
	}
	ok := cache.ValidateToken(tokenStr)
	if !ok {
		return errors.New("token is invalid")
	}
	if claims.ExpiresAt-time.Now().Unix() < config.TokenConfig.BufferTime {
		logger.DebugFmt("token is going to be expired: %s -> %s", time.Now().String(), time.Unix(claims.ExpiresAt, 0).String())
		err := DistributeToken(ctx, *claims)
		if err != nil {
			return err
		}
	}
	ctx.Set("token", tokenStr)
	ctx.Set("claims", *claims)
	ctx.Set("id", claims.UserID)
	ctx.Set("name", claims.Name)
	ctx.Set("admin", claims.Admin)
	return nil
}

func DistributeToken(ctx *gin.Context, claims TokenClaims) error {
	token := newToken()
	claims.ExpiresAt = time.Now().Unix() + config.TokenConfig.ExpiresTime
	tokenStr, err := token.createToken(claims)
	if err != nil {
		return err
	}
	err = cache.PermitToken(tokenStr, claims.UserID)
	if err != nil {
		return err
	}
	ctx.Header("New-Token", tokenStr)
	return nil
}

func RemoveToken(ctx *gin.Context) error {
	token := ctx.GetString("token")
	return cache.ExpireToken(token)
}
