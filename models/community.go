package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

func (table *Community) TableName() string {
	return "community"
}

func CreateCommunity(community Community) (int, string) {
	tx := utils.DB.Begin()
	// 事務一旦開始不論甚麼異常就一定會Rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if len(community.Name) == 0 {
		return -1, "名稱不能為空"
	}
	if community.OwnerId == 0 {
		return -1, "請先登錄"
	}

	if err := utils.DB.Create(&community).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, "建群失敗"
	}
	contact := Contact{
		OwnerId:  community.OwnerId,
		TargetId: community.ID,
		Type:     2, //群關係
	}

	if err := utils.DB.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加群失敗"
	}
	tx.Commit()
	return 200, "建群成功"
}

func LoadCommunity(ownerId uint) ([]*Community, string) {
	communities := make([]*Community, 10)
	utils.DB.Where("owner_id = ?", ownerId).Find(&communities)
	for _, v := range communities {
		if v.OwnerId == ownerId {
			return communities, ""
		}
	}
	return nil, "查無群組"
}
