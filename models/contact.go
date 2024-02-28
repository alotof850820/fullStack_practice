package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

// 人員關係
type Contact struct {
	gorm.Model
	OwnerId  uint //誰的關係信息
	TargetId uint //對應誰
	Type     int  //對應類型 1:好友 2:群聊
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriends(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)

	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id IN ?", objIds).Find(&users)

	return users
}
