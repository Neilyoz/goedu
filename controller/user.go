package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goedu/model"
	"goedu/system"
	"goedu/utils"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "failed",
			"data":    "用户名必须填写",
		})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "failed",
			"data":    "邮箱必须填写",
		})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "failed",
			"data":    "密码必须填写",
		})
		return
	}

	err := model.GetDB().Create(&user).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    "success",
	})
}

func Login(c *gin.Context) {
	type UserPost struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	// 校验传入数据是否正确
	var userPost UserPost
	_ = c.ShouldBind(&userPost)

	var user model.User
	model.GetDB().Where("email = ?", userPost.Email).Find(&user)

	if !utils.CheckPasswordHash(userPost.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "failed",
			"data":    "验证失败",
		})
		return
	}

	// 验证通过之后，生成 jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24),
	})

	data, _ := token.SignedString([]byte(system.GetConfigure().Key))

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"token": "Bearer " + data,
		},
	})
}
