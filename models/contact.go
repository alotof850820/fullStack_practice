package models

import "gorm.io/gorm"

// 人員關係
type Contact struct {
	gorm.Model
	OwnerId  uint //誰的關係信息
	TargetId uint //對應誰
	Type     int  //對應類型 0:好友 1:群聊
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
