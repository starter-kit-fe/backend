package model

// 角色菜单关联表
type UserTotp struct {
	BaseModel
	UserId uint   `json:"userId"  gorm:"comment:用户id"`
	Secret string `json:"secret"  gorm:"comment:secret"`
	URL    string `json:"url"    gorm:"comment:totp_url"`
}
