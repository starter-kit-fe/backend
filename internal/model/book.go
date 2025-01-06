package model

import (
	"time"
)

type Book struct {
	BaseModel
	Title             string    `gorm:"type:varchar(255);not null"  json:"title"`
	Author            string    `gorm:"type:varchar(255);not null"  json:"author"`
	Description       string    `gorm:"type:text;not null" json:"description"`
	Category          string    `gorm:"type:varchar(255);not null" json:"category"`
	ImgURL            string    `gorm:"type:varchar(255);not null" json:"imgUrl"`
	Status            string    `gorm:"type:varchar(50);not null" json:"status"`
	UpdateTime        time.Time `gorm:"not null" json:"updateTime"`
	LatestChapterName string    `gorm:"type:varchar(255);not null" json:"latestChapterName"`
	LatestChapterID   string    `gorm:"type:varchar(255);not null" json:"latestChapterID"`
}
