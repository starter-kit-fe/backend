package model

type BookArticle struct {
	BaseModel
	Title       string `gorm:"type:varchar(255);not null"  json:"title"`
	BookID      uint   `gorm:"type:varchar(255);not null" json:"bookId"`
	Description string `gorm:"not null" json:"description"`
	Content     string `gorm:"type:text;not null" json:"content"`
}
