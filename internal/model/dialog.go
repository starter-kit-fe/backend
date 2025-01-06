package model

type Dialog struct {
	BaseModel
	Type            *uint  `json:"type,omitempty" gorm:"comment:类型"`                 // 类型
	Model           string `json:"model,omitempty" gorm:"comment:模型"`                // 模型
	Question        string `json:"question,omitempty" gorm:"comment:问题"`             // 问题
	Answer          string `json:"answer,omitempty" gorm:"comment:回答"`               // 回答
	UserId          uint   `json:"userId" gorm:"comment:用户ID"`                       // 用户ID
	CompletionToken int    `json:"completionToken,omitempty" gorm:"comment:完成token"` // 完成token
	PromptToken     int    `json:"promptToken,omitempty" gorm:"comment:提示token"`     // 提示token
	TotalTokens     int    `json:"totalTokens,omitempty" gorm:"comment:总token"`      // 总token
}
