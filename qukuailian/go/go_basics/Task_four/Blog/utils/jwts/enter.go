package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	UserName string `json:"username"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
	UserId   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var MySecret []byte
