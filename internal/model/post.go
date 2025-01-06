package model

type Post struct {
	BaseModel
	Title   string `json:"title" gorm:"comment:標題;not null;size:255"`
	Content string `json:"content" gorm:"type:text;comment:內容"`
	Summary string `json:"summary" gorm:"type:varchar(500);comment:文章摘要"`
	Status  uint   `json:"Status" gorm:"default:null;comment:状态"`
	User    *uint  `json:"user" gorm:"comment:建立者"`
	Parent  *uint  `json:"parent,omitempty" gorm:"comment:父节点;default:null"`
}
