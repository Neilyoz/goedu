package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 首页信息
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    nil,
	})
}
