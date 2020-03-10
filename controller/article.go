package controller

import (
	"github.com/gin-gonic/gin"
	"goedu/model"
	"net/http"
	"strconv"
)

func ArticleCreate(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "认证失败，请重新认证",
			"data":    "",
		})
	}

	var article model.Article
	_ = c.ShouldBind(&article)
	userIdnew, _ := strconv.Atoi(userId.(string))
	article.UserID = uint(userIdnew)
	err := model.GetDB().Create(&article).Error
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "创建文章成功",
		"data":    "",
	})
}

func ArticleIndex(c *gin.Context) {
	var articles []model.Article
	model.GetDB().Preload("User").Order("id desc").Limit(20).Find(&articles)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    articles,
	})
}
