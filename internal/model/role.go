package model

type Role struct {
	BaseModel
	Name              string  `gorm:"column:role_name;size:30;not null" json:"role_name"`              // 角色名称
	Key               string  `gorm:"column:role_key;size:100;not null" json:"role_key"`               // 角色权限字符串
	Sort              int     `gorm:"column:role_sort;not null" json:"role_sort"`                      // 显示顺序
	Scope             string  `gorm:"column:data_scope;size:1;default:'1'" json:"data_scope"`          // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly int8    `gorm:"column:menu_check_strictly;default:1" json:"menu_check_strictly"` // 菜单树选择项是否关联显示
	DeptCheckStrictly int8    `gorm:"column:dept_check_strictly;default:1" json:"dept_check_strictly"` // 部门树选择项是否关联显示
	Status            *uint   `gorm:"column:status;size:1;default:null" json:"status"`                 // 角色状态（0正常 1停用）
	Remark            *string `gorm:"column:remark;size:500;default:null" json:"remark"`               // 备注
}
