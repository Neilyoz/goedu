package controller

import (
	"goedu/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	blackfriday "github.com/russross/blackfriday/v2"
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
	article.HtmlContent = string(blackfriday.Run([]byte(article.PlainContent)))
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

func ArticleEdit(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBind(&article)
	article.HtmlContent = string(blackfriday.Run([]byte(article.PlainContent)))
	model.GetDB().Omit("user").Save(&article)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "修改文章成功",
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

func ArticleDetail(c *gin.Context) {
	var article model.Article
	model.GetDB().Preload("User").Where("id=?", c.Param("id")).Find(&article)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    article,
	})
}
