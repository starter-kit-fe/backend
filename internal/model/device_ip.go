package model

// 设备IP关联表
type DeviceIp struct {
	BaseModel
	DeviceId uint `json:"deviceId"  gorm:"comment:设备id"`
	IPId     uint `json:"ipId"  gorm:"comment:ip id"`
}
