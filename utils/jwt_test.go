package utils

import (
	"gin/config"
	"github.com/spf13/viper"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	config.Init("../config/config.yml")
	InitJwt(viper.GetString("Jwt.Secret"), viper.GetString("Jwt.Issuer"), viper.GetInt("Jwt.ExpireDays"))
	token, err := GenerateJwtToken(&JwtClaims{
		NickName: "Rock",
		Email:    "Rock@qq.com",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)
}

func TestParseJwtToken(t *testing.T) {
	config.Init("../config/config.yml")
	InitJwt(viper.GetString("Jwt.Secret"), viper.GetString("Jwt.Issuer"), viper.GetInt("Jwt.ExpireDays"))

	claims, err := ParseJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6IlJvY2siLCJFbWFpbCI6IlJvY2tAcXEuY29tIiwiZXhwIjoxNjQ1MzY3NjE1LCJpc3MiOiJyb2Nrcm9ja3doaXRlLmNuIiwibmJmIjoxNjQ0MTU4MDE1fQ.MnTzMqlZJzZ_IDX4mF6P4EjhLIjL-P0xs2dPwB1Ox58")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(claims)
}
