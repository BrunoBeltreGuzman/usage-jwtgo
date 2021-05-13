package main

import (
	"fmt"
	"jwtgo/config"
	"jwtgo/token"

	"github.com/dgrijalva/jwt-go"
)

func main() {

	role := config.Role{
		Id:        1,
		Role:      "Admin",
		CreatedAt: "54/564/5624",
		UpdatedAt: "32/567/564",
	}

	options := config.Data{
		Id:        1,
		Name:      "Bobo",
		Email:     "bobo@gmail.com",
		Password:  "123",
		Role:      role,
		CreatedAt: "54/564/5624",
		UpdatedAt: "32/567/564",
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

	userRole := data.Role.Role
	fmt.Println(userRole)
}
