package test

import (
	"fmt"
	"jwtgo/config"
	"jwtgo/token"

	"github.com/dgrijalva/jwt-go"
)

func test() {

	role := config.Role{}

	options := config.Data{
		Id:        1,
		Name:      "Bobo",
		Email:     "Bobo",
		Password:  "Bobo",
		Role:      role,
		CreatedAt: "Bobo",
		UpdatedAt: "Bobo",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "NameApp",
		},
	}
	var createToken string = token.GetToken(options)
	fmt.Println(createToken)
	var data *config.Data = token.DecodeToken(createToken)
	name := data.Name
	fmt.Println(name)
}
