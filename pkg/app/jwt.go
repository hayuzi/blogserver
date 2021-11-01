package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hayuzi/blogserver/global"
	"time"
)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	UserType int    `json:"userType"`
	jwt.StandardClaims
}

func GetJwtSecret() string {
	return global.JWTSetting.Secret
}

func GenerateToken(userId int, username, password string, userType int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)

	claims := Claims{
		userId,
		username,
		userType,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJwtSecret())

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecret(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GetLoginClaims 获取jwt登陆用户的基础信息
func GetLoginClaims(c *gin.Context) (*Claims, error) {
	token := ""
	if s, ok := c.GetQuery("token"); ok {
		token = s
	} else {
		token = c.GetHeader("token")
	}
	var claims = &Claims{}
	claims, _ = ParseToken(token)
	return claims, nil
}
