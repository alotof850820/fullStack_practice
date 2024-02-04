package service

import (
	"ginchat/models"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用戶
// @Tags 用戶
// @Success 200 {string} json {"description":"wellcome"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	//长度为10的 models.UserBasic 指针切片
	data := make([]*models.UserBasic, 10)
	//最多10個
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 新增用戶
// @Tags 用戶
// @param name query string false "用戶名"
// @param password query string false "密碼"
// @param repassword query string false "確認密碼"
// @Success 200 {string} json {"200表示成功
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "密碼不一致",
		})
		return
	}
	user.Password = password
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用戶成功",
	})
}

// DeleteUser
// @Summary 刪除用戶
// @Tags 用戶
// @param id query string false "id"
// @Success 200 {string} json {"200表示成功
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id")) //轉type
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "刪除用戶成功",
	})
}

// UpdateUser
// @Summary 修改用戶
// @Tags 用戶
// @param id formData string false "id"
// @param name formData string false "用戶名"
// @param password formData string false "密碼"
// @param phone formData string false "手機"
// @param email formData string false "郵箱"
// @Success 200 {string} json {"200表示成功
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id")) //轉type
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(-1, gin.H{
			"message": err.Error(),
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "修改用戶成功",
		})
	}

}
