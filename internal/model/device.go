package model

type Device struct {
	BaseModel
	Os             string `json:"os" gorm:"comment:系统名称"`
	OsVersion      string `json:"osVersion" gorm:"comment:设备版本"`
	Type           string `json:"type" gorm:"comment:设备类型"`
	Finger         string `json:"finger" gorm:"primaryKey;index:idx_device_finger;unique;comment:设备指纹"`
	BrowserName    string `json:"browse_name" gorm:"comment:浏览器名称"`
	BrowserVersion string `json:"browse_version" gorm:"comment:浏览器名称"`
	UserAgent      string `json:"user_agent" gorm:"comment:特征信息"`
}
