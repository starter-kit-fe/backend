package model

// 用户设备关联表
type UserDevice struct {
	BaseModel
	UserId   uint `json:"userId"  gorm:"comment:用户id"`
	DeviceId uint `json:"deviceId"  gorm:"comment:角色id"`
}
