package jwtx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
)

var CtxKeyJwtUid = "jwtUserId" //用户端

func GetToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[CtxKeyJwtUid] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetUid(ctx context.Context) int64 {
	uid, _ := ctx.Value(CtxKeyJwtUid).(json.Number).Int64()
	return uid
}

func GetUidByToken(secretKey string, tokenString string) int64 {

	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	}); err == nil {
		Valid := token.Valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && Valid {

			if uid, err := strconv.Atoi(fmt.Sprint(claims[CtxKeyJwtUid])); err == nil {
				return int64(uid)
			}
		}
	}

	return 0
}
