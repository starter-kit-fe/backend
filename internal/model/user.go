package model

import (
	"time"
)

type User struct {
	BaseModel
	Birthday time.Time `json:"birthday,omitempty" gorm:"type:date;default:null;comment:生日"`
	UUID     string    `json:"-" gorm:"primaryKey;type:char(36);index:idx_uuid,unique;comment:用户UUID"`
	Email    string    `json:"email" gorm:"uniqueIndex:idx_email;type:varchar(50);comment:用户邮箱"`
	Password string    `json:"-" gorm:"type:char(60);comment:用户登录密码（哈希存储）" `
	NickName string    `json:"nickName" gorm:"type:varchar(32);comment:用户昵称"`
	Avatar   string    `json:"avatar" gorm:"type:varchar(255);comment:用户头像"`
	Phone    string    `json:"phone,omitempty" gorm:"type:varchar(20);index:idx_phone;comment:用户手机号"`
	Gender   uint      `json:"gender,omitempty" gorm:"type:smallint;comment:用户性别"`
	Inviter  uint      `json:"inviter,omitempty" gorm:"comment:邀请用户ID"`
	Remark   string    `json:"remark,omitempty" gorm:"size:500;comment:备注"`
}
