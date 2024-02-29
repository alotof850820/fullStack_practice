package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// FindUserByNameAndPassword
// @Summary 登入
// @Tags 用戶
// @param name query string false "用戶名"
// @param password query string false "密碼"
// @Success 200 {string} json {"code":"message"}
// @Router /user/findUserByNameAndPassword [post]
func FindUserByNameAndPassword(c *gin.Context) {
	data := models.UserBasic{}
	// name := c.Query("name") //獲取 URL Query Parameters 中的值。
	// password := c.Query("password")
	name := c.Request.FormValue("name") //獲取 POST 請求的 Form Data 中的值。 application/x-www-form-urlencoded
	password := c.Request.FormValue("password")

	// 先查人
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "查無此用戶",
			"data":    user,
		})
		return
	}

	// 再查密碼
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "密碼錯誤",
			"data":    user,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPassword(name, pwd)

	//拿用戶信息 驗證token
	user.ID = data.ID
	// user.Identity

	c.JSON(200, gin.H{
		"code":    200, // -1 失敗
		"message": "登入成功",
		"data":    data,
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
	// user.Name = c.Query("name")
	// password := c.Query("password")
	// repassword := c.Query("repassword")
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Name)
	if data.Name == "" || password == "" || repassword == "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用戶名子或密碼不得為空",
			"data":    user,
		})
		return
	}

	if data.Name != "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用戶名子已存在",
			"data":    user,
		})
		return
	}

	if password != repassword {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "密碼不一致",
			"data":    user,
		})
		return
	}
	// 生成31位随机整数，并一定格式化为六位数字串。
	// 推荐使用crypto/rand
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt

	result := models.CreateUser(user)
	if result.Error != nil {
		// 处理错误
		fmt.Println("Error creating user:", result.Error)
		c.JSON(-1, gin.H{
			"message": "新增用戶失敗",
			"error":   result.Error.Error(),
			"data":    user,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "新增用戶成功",
		"data":    user,
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
		"code":    200,
		"message": "刪除用戶成功",
		"data":    user,
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
			"code":    -1,
			"message": err.Error(),
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    200,
			"message": "修改用戶成功",
			"data":    user,
		})
	}
}

// 防止跨域站點偽造請求
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println("close:", err)
		}
	}(ws) // 将 ws 传递给自己

	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println("subscribe:", err)
		return
	}
	tm := time.Now().Format("2006-01-02 15:04:05")
	m := fmt.Sprintf("[ws][%s] %s", tm, msg)
	err = ws.WriteMessage(1, []byte(m))
	if err != nil {
		fmt.Println("write:", err)
		return
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}

// SearchFriends godoc
// @Summary 搜尋朋友
// @Description 通過用戶 ID 搜尋朋友
// @Tags 用戶
// @Accept json
// @Produce json
// @Param userId formData int true "用戶 ID"
// @Router /searchFriends [post]
func SearchFriends(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriends(uint(userId))
	utils.RespOKList(c.Writer, users, len(users))
}

func AddFriend(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.FormValue("userId"))
	targetId, _ := strconv.Atoi(c.Request.FormValue("targetId"))
	code, msg := models.AddFriend(uint(userId), uint(targetId))
	if code == 200 {
		utils.RespOK(c.Writer, msg, code)
	} else {
		utils.RespFail(c.Writer, msg)
	}

}
