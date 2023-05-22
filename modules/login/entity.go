package login

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id			int		`json:"id"`
	Fullname	string	`json:"fullname"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type UserResponse struct {
	Token *MyClaims	`json:"token"`
}

var APPLICATION_NAME = "My Go App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNATURE_KEY = []byte("the hashslinging slyser")

type MyClaims struct {
	jwt.StandardClaims
	Idp			int 	`json:"id_user"`
	Fullname 	string 	`json:"fullname"`
}