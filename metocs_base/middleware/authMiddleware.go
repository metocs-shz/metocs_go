package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"metocs/base/config"
	"net/http"
	"strings"
	"time"
)

type UserClaims struct {
	ID     string `json:"ID"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Status int    `json:"status"`
	//jwt-go提供的标准claim
	jwt.RegisteredClaims
}

func GenerateToken(claims *UserClaims) string {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().AddDate(0, 0, config.App.Auth.Expires))
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.App.Auth.Secret))
	if err != nil {
		log.Println("身份令牌生成失败！----------> ", err.Error())
	}
	return sign
}

//Authorize 授权中间件
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未授权访问！"})
		}
		authorization = strings.Replace(authorization, "Bearer ", "", -1)
		token := parseToken(authorization)
		c.Set("UserClaims", token)
		c.Next()
	}
}

func parseToken(tokenString string) *UserClaims {
	//解析token

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.Auth.Secret), nil
	})

	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		log.Println("令牌解析失败 ---------------> ", err.Error())
	}
	return claims
}
