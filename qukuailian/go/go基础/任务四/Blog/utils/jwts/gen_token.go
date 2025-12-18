package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"modulename/global"
	"time"
)

// 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
