package service

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// GetVueHtml
// @Tags 首頁
// @Success 200 {string} wellcome
// @Router /index [get]
func GetVueHtml(c *gin.Context) {
	ind := template.Must(template.ParseFiles("./views/index.html"))
	ind.Execute(c.Writer, "index")
}
