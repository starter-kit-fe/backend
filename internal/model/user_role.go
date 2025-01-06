package model

// 角色菜单关联表
type UserRole struct {
	BaseModel
	UserId uint `json:"userId"  gorm:"comment:用户id"`
	RoleId uint `json:"roleId"  gorm:"comment:角色id"`
}
