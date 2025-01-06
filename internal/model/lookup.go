package model

type Lookup struct {
	BaseModel
	GroupValue string `gorm:"type:varchar(100);not null;index:idx_group_value_label" json:"groupValue" ` // 类型
	EntryLabel string `gorm:"type:varchar(100);not null" json:"entryLabel" `                             // 常量值
	EntryValue string `gorm:"type:varchar(255);primaryKey;not null;uniqueIndex" json:"entryValue" `      // 唯一标识符
	SortOrder  uint   `gorm:"default:0;" json:"sortOrder"`                                               // 排序标记
	Status     uint   `gorm:"type:smallint;default:1;" json:"status" `                                   // 状态
	Remark     string `gorm:"size:500;default:''" json:"remark"`                                         // 备注
}
