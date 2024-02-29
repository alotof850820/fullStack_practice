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

func AddFriend(ownerId uint, targetId uint) (int, string) {
	user := UserBasic{}
	if targetId != 0 {
		user = FindByID(targetId)
		if user.Identity != "" {
			if ownerId == targetId {
				return -1, "不能添加自己"
			}
			contact0 := Contact{}
			utils.DB.Where("owner_id = ? and target_id = ? and type = 1", ownerId, targetId).First(&contact0)
			if contact0.ID != 0 {
				return -1, "不能重複添加"
			}

			tx := utils.DB.Begin() // 创数据库事务。保证数据的完整性。
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback() //recover用於 从 panic 中恢复 並回朔
				}
			}()
			contact := Contact{
				OwnerId:  ownerId,
				TargetId: targetId,
				Type:     1,
			}
			// 對方也要創建
			contact1 := Contact{
				OwnerId:  targetId,
				TargetId: ownerId,
				Type:     1,
			}
			if err := utils.DB.Create(&contact).Error; err != nil {
				tx.Rollback() // 撤销之前的所有数据库操作，确保数据的一致性。
				return -1, "添加失敗"
			}
			if err := utils.DB.Create(&contact1).Error; err != nil {
				tx.Rollback()
				return -1, "添加失敗"
			}

			tx.Commit()
			return 200, "添加成功"
		}
		return -1, "用戶不存在"
	}
	return -1, "Id 不能為空"
}
