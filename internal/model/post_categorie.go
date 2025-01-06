package model

// 分类文章关联表
type PostCategorie struct {
	BaseModel
	PostId     uint `json:"postId" gorm:"comment:文章id"`
	CategoryId uint `json:"categoryId" gorm:"comment:分类id"`
}
