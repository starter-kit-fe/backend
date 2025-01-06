package model

// 角色菜单关联表
type RoleMenu struct {
	BaseModel
	MenuId uint `json:"menuId"  gorm:"comment:菜单id"`
	RoleId uint `json:"roleId"  gorm:"comment:角色id"`
}
