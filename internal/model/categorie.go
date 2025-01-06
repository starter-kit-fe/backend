package model

type Categorie struct {
	BaseModel
	// Category is the category of the knowledge and cannot be null. It has a maximum length of 255 characters.
	Name   string `json:"name" gorm:"not null;unique;comment:分类名称"` // 分类名称不能为空，并且在数据库中必须是唯一的
	Remark string `json:"remark" gorm:"size:255;comment:分类描述"`      // 可选的分类描述，限制大小为255字符
}
