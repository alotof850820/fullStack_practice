package sql

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Identity      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)~請輸入正確的手機號碼"`
	Email         string `valid:"email~請輸入正確的郵箱格式"`
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LoginOutTime  *time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
	Avatar        string
}

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

type Message struct {
	gorm.Model
	UserId     int64  //發送者
	TargetId   int64  //接收者
	Type       int    //消息類型 群聊 私聊 廣播
	Media      int    //消息類型 文字 圖片 音訊
	Content    string //消息內容
	Url        string //文件地址
	Pic        string //圖片地址
	Desc       string //描述
	Amount     int    //其他數字統計
	CreateTime uint64
	ReadTime   uint64 //讀取時間
}

type Contact struct {
	gorm.Model
	OwnerId  uint //誰的關係信息
	TargetId uint //對應誰
	Type     int  //對應類型 1:好友 2:群聊
	Desc     string
}

type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

// 可重新run修改table
func CreateTable() {
	// db := utils.DB

	// // 自动生成相应的数据库表，并确保表的结构与模型的结构一致。
	// db.AutoMigrate(&models.Community{})
	// db.AutoMigrate(&models.UserBasic{})
	// db.AutoMigrate(&models.Message{})
	// db.AutoMigrate(&models.Contact{})
	// db.AutoMigrate(&models.GroupBasic{})

	// gorm創建
	// user := &models.UserBasic{}   // 创建一个 models.UserBasic 结构体的实例
	// user.Name = "申專"            // 设置该实例的 Name 属性为 "申專"
	// db.Create(user)               // 使用 Gorm 将该实例插入到数据库

	// gorm查询数据库
	// fmt.Println(db.First(user, 1)) // 查主键为 1，并存储到 user 中。

	// gorm更新数据库
	// db.Model(user).Update("Password", "1234") // 更新 user 的 Password 属性为 "1234"
}
