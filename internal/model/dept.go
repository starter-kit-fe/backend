package model

// SysDept 部门表模型
// DeptID 部门id
// ParentID 父部门id
// Ancestors 祖级列表
// DeptName 部门名称
// Sort 显示顺序
// Leader 负责人
// Phone 联系电话
// Email 邮箱
// Status 部门状态（0正常 1停用）
// CreateBy 创建者
// UpdateBy 更新者
type Dept struct {
	BaseModel
	DeptID    uint    `gorm:"primaryKey;autoIncrement;column:dept_id" json:"deptId"`
	ParentID  uint    `gorm:"default:0;column:parent_id" json:"parentId"`
	Ancestors string  `gorm:"size:50;default:'';column:ancestors" json:"ancestors"`
	DeptName  string  `gorm:"size:30;default:'';column:dept_name" json:"deptName"`
	Sort      int     `gorm:"default:0;column:sort" json:"Sort"`
	Leader    *string `gorm:"size:20;default:null;column:leader" json:"leader"`
	Phone     *string `gorm:"size:11;default:null;column:phone" json:"phone"`
	Email     *string `gorm:"size:50;default:null;column:email" json:"email"`
	Status    *uint   `gorm:"size:1;default:null;column:status" json:"status"`
}
