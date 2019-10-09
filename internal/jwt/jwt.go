package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secret = []byte("this is a test secret")
)

type Msg struct {
	UserID   int32  `json:"userid"`
	UserName string `json:"username"`
}

type MyCustomClaims struct {
	Msg
	jwt.StandardClaims
}

func Encode(r Msg) (string, error) {
	claims := MyCustomClaims{
		r,
		jwt.StandardClaims{
			// ExpiresAt: time.Now().Unix(),// 过期时间
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), //  过 期时间
			Issuer:    "com.example",                         //  该JWT的签发者,可选
			// Subject:"",// 该JWT所面向的用户
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func Decode(tokenString string) (Msg, error) {
	//  Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return Msg{}, err
	}
	//  Validate the token and return the custom claims
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims.Msg, nil
	}
	return Msg{}, err
}
