package config

import (
	"github.com/dgrijalva/jwt-go"
)

//data
type Data struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      Role   `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	jwt.StandardClaims
}

type Role struct {
	Id        int    `json:"id"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const KEY string = "secureSecretText"
