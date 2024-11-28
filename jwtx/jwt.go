package jwtx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

type JwtClient struct {
	ctx    context.Context
	ctxKey string
}

func NewJwtX(context context.Context, key string) *JwtClient {
	return &JwtClient{
		ctxKey: key,
		ctx:    context,
	}
}

func (j *JwtClient) GetToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[j.ctxKey] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (j *JwtClient) GetUid() int64 {
	var uid int64
	if jsonUid, ok := j.ctx.Value(j.ctxKey).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		}
	}
	return uid
}

func (j *JwtClient) GetUidByToken(secretKey string, tokenString string) int64 {

	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	}); err == nil {
		Valid := token.Valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && Valid {
			if uid, err := strconv.Atoi(fmt.Sprint(claims[j.ctxKey])); err == nil {
				return int64(uid)
			}
		}
	}
	return 0
}
