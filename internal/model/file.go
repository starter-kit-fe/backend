package model

type File struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Hash string `gorm:"uniqueIndex;type:varchar(255);not null" json:"hash"`
	Size int64  `gorm:"type:bigint;not null" json:"size"`
	Url  string `gorm:"type:varchar(255);" json:"url"`
}
