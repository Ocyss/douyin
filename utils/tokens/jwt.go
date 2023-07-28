package tokens

import (
	"errors"
	"github.com/Ocyss/douyin/internal/conf"
	"github.com/golang-jwt/jwt"
	"time"
)

var JwtKey = []byte(conf.Conf.JwtSecret)

type MyClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GetToken 生成token
func GetToken(id int64, username string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 3)
	SetClaims := MyClaims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ByteHunters",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	return reqClaim.SignedString(JwtKey)
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	setToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, nil
	} else {
		return nil, errors.New("验证失败")
	}
}
