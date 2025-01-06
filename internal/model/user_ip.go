package model

// 用户IP关联表
type UserIp struct {
	BaseModel
	UserId uint `json:"userId"  gorm:"comment:用户id"`
	IPId   uint `json:"ipId"  gorm:"comment:ip id"`
}
