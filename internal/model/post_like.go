package model

// 文章点赞表
type PostLike struct {
	BaseModel
	PostId    uint  `json:"postId" gorm:"comment:文章id"`
	UserID    uint  `json:"userId" gorm:"comment:用户id"`
	LikeCount *uint `json:"view_count" gorm:"comment:点赞次数"`
}
