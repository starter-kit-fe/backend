package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedBy *uint `gorm:"column:create_by;size:64;" json:"created_by"` // 创建者
	UpdatedBy *uint `gorm:"column:update_by;size:64;" json:"updated_by"`
}

func getUserIDFromContext(tx *gorm.DB) *uint {
	userId, ok := tx.Statement.Context.Value("userId").(*uint) // 假设你有一个 User 类型
	println(userId)
	if ok && userId != nil {
		return userId // 返回用户的IDd
	}
	return nil
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	userID := getUserIDFromContext(tx) // 从上下文中获取用户ID
	if userID != nil {
		tx.Statement.SetColumn("update_by", userID)
		tx.Statement.SetColumn("created_by", userID)
	}
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	userID := getUserIDFromContext(tx) // 从上下文中获取用户ID
	if userID != nil {
		m.UpdatedBy = userID
		tx.Statement.SetColumn("update_by", userID)
	}
	return
}
