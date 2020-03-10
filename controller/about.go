package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"goedu/model"
	"net/http"
)

func AboutDetail(c *gin.Context) {
	var page model.Page
	slug := c.Param("slug")
	model.GetDB().Where("slug = ?", slug).First(&page)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取成功",
		"data":    page,
	})
}

func AboutStore(c *gin.Context) {
	var page model.Page

	_ = c.ShouldBind(&page)
	page.Slug = "about"
	page.HtmlContent = string(blackfriday.Run([]byte(page.PlainContent)))

	// 新建的时候
	if page.ID == 0 {
		err := model.GetDB().Create(&page).Error
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "修改成功",
			"data":    "",
		})
	} else {
		model.GetDB().Save(&page)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "修改成功",
			"data":    "",
		})
	}
}
