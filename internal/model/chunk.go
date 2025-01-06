package model

type Chunk struct {
	BaseModel
	FileHash string `gorm:"type:varchar(255);not null" json:"file_hash"`
	Hash     string `gorm:"uniqueIndex;type:varchar(255);not null" json:"hash"`
	Index    int64  `gorm:"type:bigint;not null" json:"index"`
	Total    int64  `gorm:"type:bigint;not null" json:"total"`
}
