package utils

import (
	"github.com/RockRockWhite/LabWeb.API/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtClaims struct {
	Id          uint   // 用户id
	Username    string // 昵称
	Fullname    string // 实验室成员全名
	Email       string // 邮箱
	VerifyState bool   // 邮箱验证状态
	Telephone   string // 手机号码
	IsAdmin     bool   // 是否管理员
	jwt.StandardClaims
}

var _secret []byte
var _issuer string
var _expireDays int

func init() {
	_secret = []byte(config.GetString("Jwt.Secret"))
	_issuer = config.GetString("Jwt.Issuer")
	_expireDays = config.GetInt("Jwt.ExpireDays")
}

// GenerateJwtToken 生成JwtToken
func GenerateJwtToken(claims *JwtClaims) (string, error) {
	claims.Issuer = _issuer
	claims.NotBefore = int64(time.Now().Unix())
	claims.ExpiresAt = int64(time.Now().AddDate(0, 0, _expireDays).Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(_secret)
}

// ParseJwtToken 解码JwtToken
func ParseJwtToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return _secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
