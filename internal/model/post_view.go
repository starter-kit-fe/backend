package model

// 角色菜单关联表
type PostView struct {
	BaseModel
	PostId    uint `json:"postId" gorm:"comment:文章id"`
	ViewCount uint `json:"view_count" gorm:"comment:浏览次数"`
}
