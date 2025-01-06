package model

// V2Log represents the log table in the database.
type Ip struct {
	BaseModel
	Ip          string   `json:"ip" gorm:"primaryKey;index:idx_ip;unique;comment:ip地址"`
	Lat         *float32 `json:"lat" gorm:"comment:坐标"`
	Lon         *float32 `json:"lon" gorm:"comment:坐标"`
	Country     *string  `json:"country" gorm:"size:100;comment:国家名称"`
	CountryCode *string  `json:"countryCode" gorm:"size:100;comment:国家代码"`
	Region      *string  `json:"region" gorm:"comment:省份"`
	RegionName  *string  `json:"regionName" gorm:"comment:省份名称"`
	City        *string  `json:"city" gorm:"comment:城市"`
	Zip         *string  `json:"zip" gorm:"comment:邮政编码"`
	Isp         *string  `json:"isp" gorm:"comment:运营商"`
	Org         *string  `json:"org" gorm:"comment:组织"`
	As          *string  `json:"as" gorm:"comment:as"`
}
