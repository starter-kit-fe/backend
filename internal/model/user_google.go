package model

// 角色菜单关联表
type UserGoogle struct {
	BaseModel
	UserId   uint   `json:"userId"  gorm:"comment:用户id"`
	GoogleId string `json:"googleId"  gorm:"comment:id"`
}
