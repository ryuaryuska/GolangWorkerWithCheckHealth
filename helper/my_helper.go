
package helper

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
	"WorkerWithCheckHealth/config"
	"WorkerWithCheckHealth/exception"
	"WorkerWithCheckHealth/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

func InsertRedis(payload model.SetDataRedis) {
	var redis = config.RedisConnection()
	json, err := json.Marshal(payload.Data)
	exception.PanicIfNeeded(err)
	err = redis.Set(payload.Key, json, payload.Exp).Err()
	exception.PanicIfNeeded(err)
}

func GetRedis(payload model.Redis) interface{} {
	var redis = config.RedisConnection()

	value, _ := redis.Get(payload.Key).Result()
	return value
}

func DelRedis(payload model.Redis) {
	var redis = config.RedisConnection()
	redis.Del(payload.Key)
}

func CreateToken(request model.JwtPayload) *model.TokenDetails {
	accessExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_MINUTE"))
	refreshExpired, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_DAYS"))

	td := &model.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessExpired)).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * time.Duration(refreshExpired)).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error

	keyAccess, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_ACCESS_PRIVATE_KEY")))
	exception.PanicIfNeeded(err)

	keyRefresh, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_REFRESH_PRIVATE_KEY")))
	exception.PanicIfNeeded(err)
	now := time.Now().UTC()

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = request.UserID
	atClaims["username"] = request.Username
	atClaims["role_id"] = request.RoleID
	atClaims["role_name"] = request.RoleName
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["exp"] = td.AtExpires
	atClaims["iat"] = now.Unix()
	atClaims["iss"] = "topindopay"
	atClaims["aud"] = "topindopay"

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at.Header["topindopay"] = "jwt"
	td.AccessToken, err = at.SignedString(keyAccess)

	if err != nil {
		exception.PanicIfNeeded(errors.New(model.AUTHENTICATION_FAILURE_ERR_TYPE))
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = request.UserID
	rtClaims["username"] = request.Username
	rtClaims["role_id"] = request.RoleID
	rtClaims["role_name"] = request.RoleName
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodRS256, rtClaims)

	td.RefreshToken, err = rt.SignedString(keyRefresh)
	if err != nil {
		exception.PanicIfNeeded(errors.New(model.AUTHENTICATION_FAILURE_ERR_TYPE))
	}

	return td

}

func CreateAuth(request model.JwtPayload, td *model.TokenDetails) {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	InsertRedis(model.SetDataRedis{
		Key:  td.AccessUuid,
		Data: request,
		Exp:  at.Sub(now),
	})

	InsertRedis(model.SetDataRedis{
		Key:  td.RefreshUuid,
		Data: request,
		Exp:  rt.Sub(now),
	})

}
