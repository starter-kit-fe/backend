package model

type Api struct {
	BaseModel
	// Category is the category of the knowledge and cannot be null. It has a maximum length of 255 characters.
	// Method 表示 HTTP 请求方法，例如 GET、POST、PUT、DELETE 等
	Method string `gorm:"type:varchar(10);not null" json:"method"`

	// Path 表示路由的路径，例如 /example
	Path string `gorm:"type:varchar(255);not null" json:"path"`

	// Handler 表示处理该路由请求的函数名
	Handler string `gorm:"type:varchar(255);not null" json:"handler"`
}
