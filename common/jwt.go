package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kaijian/gin-vue/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//过期时间
	expirationTime := time.Now().Add(7 * 27 * time.Hour).Unix()
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			//发放时间
			IssuedAt: time.Now().Unix(),
			Issuer:   "oceanLearn.tech", //谁发放
			Subject:  "user token",      //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString,nil
}

func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
	claims := &Claims{}

	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims,err
}
