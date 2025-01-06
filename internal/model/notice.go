package model

type Notice struct {
	BaseModel
	Title   string  `gorm:"column:title;size:50;not null" json:"title"`            // 公告标题
	Type    *uint   `gorm:"column:type;size:1;default:null" json:"type,omitempty"` // 公告类型（1通知 2公告）`
	Content string  `gorm:"column:content;type:text;default:null" json:"content"`  // 公告内容
	Status  *uint   `gorm:"column:status;size:1;default:null" json:"status"`       // 公告状态（0正常 1关闭）
	Remark  *string `gorm:"column:remark;size:255;default:null" json:"remark"`     // 备注
}
