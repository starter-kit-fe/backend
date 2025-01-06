package model

import "time"

// 短链生成
type ShortLink struct {
	BaseModel
	OriginalURL string     `json:"originalUrl"  gorm:"comment:原始链接"`
	ShortCode   string     `json:"shortCode"  gorm:"comment:短链接"`
	UserId      uint       `json:"userId"  gorm:"comment:用户id"`
	ExpiresAt   *time.Time `json:"expiresAt"  gorm:"comment:过期时间"`
	ClickCount  uint       `json:"clickCount"  gorm:"comment:点击次数"`
}
