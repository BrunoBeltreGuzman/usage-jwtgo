package token

import (
	"github.com/dgrijalva/jwt-go"
	"jwtgo/config"
	"log"
)

func DecodeToken(token string) *config.Data {
	//Valid Token
	validToken, err := jwt.ParseWithClaims(
		token,
		&config.Data{},
		func(validToken *jwt.Token) (interface{}, error) {
			return []byte(config.KEY), nil
		},
	)
	data, ok := validToken.Claims.(*config.Data)
	if !ok {
		log.Fatal(err)
	}

	//time
	//fmt.Println(data.ExpiresAt)

	//get user
	return data
}

func GetToken(options config.Data) string {
	objectJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, options)
	token, err := objectJWT.SignedString([]byte(config.KEY))
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	//Token
	return token
}
