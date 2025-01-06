package model

type Permissions struct {
	BaseModel
	Name     string `gorm:"size:50;not null" json:"name"`                  // 菜单名称
	Type     uint   `gorm:"type:smallint;default:0" json:"type"`           // 菜单类型
	ParentID uint   `gorm:"type:smallint;default:0;index" json:"parentId"` // 父菜单ID
	Sort     uint   `gorm:"default:0" json:"sort"`                         // 显示顺序
	Path     string `gorm:"size:200;default:''" json:"path"`               // 路由地址
	Perms    string `gorm:"primaryKey;size:100;unique" json:"perms"`       // 权限标识
	Icon     string `gorm:"size:100;default:''" json:"icon"`               // 菜单图标
	Remark   string `gorm:"size:500;default:''" json:"remark"`             // 备注
	IsFrame  uint   `gorm:"type:smallint;default:0" json:"isFrame"`
	Status   uint   `gorm:"type:smallint;default:0" json:"status"`
}
