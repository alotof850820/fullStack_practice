package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 可重新run修改table
func main() {
	dsn := "root:Oo38560067@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动生成相应的数据库表，并确保表的结构与模型的结构一致。
	db.AutoMigrate(&models.Community{})
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
