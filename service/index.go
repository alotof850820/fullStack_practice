package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Tags 首頁
// @Success 200 {string} wellcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "wellcome",
	})
}
