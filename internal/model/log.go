package model

type Log struct {
	BaseModel
	// UserId is the user id of the log.
	UserId *uint `json:"userId,omitempty" gorm:"comment:用户id"`
	// Method is the method of the log and cannot be null.
	Method string `json:"method" gorm:"not null;comment:请求方法"`
	// 接口请求花费时间 单位毫秒
	CostTime int64 `json:"costTime"  gorm:"not null;comment:接口请求花费时间单位毫秒(ms)"`
	// ResCode is the response code of the log.
	Code int `json:"code"  gorm:"not null;comment:接口请求状态码"`
	// ResSize is the response size of the log.
	Size int `json:"size"  gorm:"not null;comment:接口请求返回大小"`
	// IP is the IP address of the log.
	Path string `json:"path" gorm:"size:255"`
	// Query is the query of the log.
	Query string `json:"query" gorm:"size:255"`
	// Body is the body of the log.
	Body *string `json:"body" gorm:"type:text"`
	// IP is the IP address of the log.
	IpId *uint `json:"ipId,omitempty"  gorm:"comment:ip地址"`
	// DeviceId is the device id of the log.
	DeviceId *uint `json:"deviceId,omitempty"  gorm:"comment:设备id"`
}
