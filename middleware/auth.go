package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goedu/system"
	"net/http"
	"strings"
)

// parseToken 校验Token
func parseToken(token string) (string, error)  {
	parseAuth, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(system.GetConfigure().Key), nil
	})

	if err != nil {
		return "", errors.New("Token Error!!!")
	}

	claims := parseAuth.Claims.(jwt.MapClaims)

	return fmt.Sprintf("%d", int(claims["user_id"].(float64))), nil
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := gin.H{
			"code":    http.StatusUnauthorized,
			"message": "无法认证，请重新登录",
			"data":    nil,
		}

		auth := context.Request.Header.Get("Authorization")

		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result)
			return
		}

		auth = strings.Split(auth, " ")[1]

		// 校验token
		userId, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result["message"] = "Error Message：" + err.Error()
			context.JSON(http.StatusUnauthorized, result)
		} else {
			println("token 正确")
		}
		context.Set("user_id", userId)

		context.Next()
	}
}
