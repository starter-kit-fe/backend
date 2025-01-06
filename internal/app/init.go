package app

import (
	"admin/internal/model"
	"log"
)

func (a *App) Init() {
	a.migrate()
	a.initApi()
	a.init()

}
func (a *App) initApi() {
	routes := a.router.Routes()
	for _, route := range routes {
		if err := a.db.FirstOrCreate(&model.Api{
			Path:    route.Path,
			Method:  route.Method,
			Handler: route.Handler,
		}, &model.Api{
			Path: route.Path,
		}).Error; err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) migrate() {
	a.db.AutoMigrate(
		&model.Api{},
		&model.BookArticle{},
		&model.Book{},
		&model.Categorie{},
		&model.Chunk{},
		&model.Dept{},
		&model.DeviceIp{},
		&model.Device{},
		&model.Dialog{},
		&model.Lookup{},
		&model.File{},
		&model.Ip{},
		&model.Log{},
		&model.Notice{},
		&model.Permissions{},
		&model.PostCategorie{},
		&model.PostLike{},
		&model.PostView{},
		&model.Post{},
		&model.RoleMenu{},
		&model.Role{},
		&model.UserDevice{},
		&model.UserGoogle{},
		&model.UserIp{},
		&model.UserRole{},
		&model.User{},
		&model.UserTotp{},
	)
}

func (a *App) init() {
	lookups := []*model.Lookup{
		{
			GroupValue: "状态",
			EntryValue: "active",
			EntryLabel: "活跃/启用",
			Remark:     "",
			Status:     1,
			SortOrder:  1,
		},
		{

			GroupValue: "状态",
			EntryValue: "inactive",
			EntryLabel: "不活跃/禁用",
			Remark:     "",
			Status:     1,
			SortOrder:  2,
		},
		{

			GroupValue: "状态",
			EntryValue: "checked",
			EntryLabel: "选中",
			Remark:     "",
			Status:     1,
			SortOrder:  3,
		},
		{

			GroupValue: "状态",
			EntryValue: "approved",
			EntryLabel: "已批准",
			Remark:     "",
			Status:     1,
			SortOrder:  4,
		},
		{

			GroupValue: "状态",
			EntryValue: "rejected",
			EntryLabel: "已拒绝",
			Remark:     "",
			Status:     1,
			SortOrder:  5,
		},
		{

			GroupValue: "状态",
			EntryValue: "completed",
			EntryLabel: "完成",
			Remark:     "",
			Status:     1,
			SortOrder:  6,
		},
		{

			GroupValue: "状态",
			EntryValue: "canceled",
			EntryLabel: "取消",
			Remark:     "",
			Status:     1,
			SortOrder:  7,
		},
		{

			GroupValue: "状态",
			EntryValue: "archived",
			EntryLabel: "归档",
			Remark:     "",
			Status:     1,
			SortOrder:  8,
		},
		{

			GroupValue: "状态",
			EntryValue: "deleted",
			EntryLabel: "删除（逻辑删除）",
			Remark:     "",
			Status:     1,
			SortOrder:  9,
		},
		{

			GroupValue: "状态",
			EntryValue: "draft",
			EntryLabel: "草稿",
			Remark:     "",
			Status:     1,
			SortOrder:  10,
		},
		{

			GroupValue: "性别",
			EntryValue: "male",
			EntryLabel: "男",
			Remark:     "",
			Status:     1,
			SortOrder:  1,
		},
		{

			GroupValue: "性别",
			EntryValue: "female",
			EntryLabel: "女",
			Remark:     "",
			Status:     1,
			SortOrder:  2,
		},
		{

			GroupValue: "权限类型",
			EntryValue: "M",
			EntryLabel: "菜单",
			Remark:     "权限类型",
			Status:     1,
			SortOrder:  1,
		},
		{

			GroupValue: "权限类型",
			EntryValue: "F",
			EntryLabel: "目录",
			Remark:     "权限类型",
			Status:     1,
			SortOrder:  2,
		},
		{

			GroupValue: "权限类型",
			EntryValue: "B",
			EntryLabel: "按钮",
			Remark:     "权限类型",
			Status:     1,
			SortOrder:  3,
		},
		{

			GroupValue: "请求类型",
			EntryLabel: "GET",
			EntryValue: "GET",
			Remark:     "请求类型",
			Status:     1,
			SortOrder:  1,
		}, {

			GroupValue: "请求类型",
			EntryLabel: "POST",
			EntryValue: "POST",
			Remark:     "请求类型",
			Status:     1,
			SortOrder:  2,
		},
		{

			GroupValue: "请求类型",
			EntryLabel: "DELETE",
			EntryValue: "DELETE",
			Remark:     "请求类型",
			Status:     1,
			SortOrder:  3,
		},
		{

			GroupValue: "请求类型",
			EntryLabel: "PUT",
			EntryValue: "PUT",
			Remark:     "请求类型",
			Status:     1,
			SortOrder:  4,
		}, {

			GroupValue: "请求类型",
			EntryLabel: "PATCH",
			EntryValue: "PATCH",
			Remark:     "请求类型",
			Status:     1,
			SortOrder:  5,
		},
	}

	for _, it := range lookups {
		if err := a.repos.Lookup.Create(it); err != nil {
			log.Printf("failed to initialize lookup data: %v", err)
		}
	}
}
