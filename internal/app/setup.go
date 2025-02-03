package app

import (
	"admin/internal/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func (a *App) Setup() {
	// 检查数据库是否需要初始化
	if err := a.checkAndMigrateIfNeeded(); err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
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

func (a *App) checkAndMigrateIfNeeded() error {
	// 尝试检查表是否存在
	tableExists := a.db.Migrator().HasTable(&model.User{})

	if !tableExists {
		log.Println("Tables don't exist, performing initial migration...")
		// 执行迁移
		a.migrate()
		// 执行初始化
		a.initApi()
		a.init()
		log.Println("Initial setup completed successfully")
		return nil
	}

	// 如果表存在，检查是否需要初始化数据
	needsInit, err := a.needsInitialization()
	if err != nil {
		return fmt.Errorf("failed to check initialization status: %v", err)
	}

	if needsInit {
		log.Println("Tables exist but data initialization needed...")
		a.initApi()
		a.init()
		log.Println("Data initialization completed successfully")
	} else {
		log.Println("Database already initialized, skipping setup")
	}

	return nil
}

func (a *App) needsInitialization() (bool, error) {
	var apiCount, lookupCount int64

	if err := a.db.Model(&model.Api{}).Count(&apiCount).Error; err != nil {
		return false, fmt.Errorf("failed to count APIs: %v", err)
	}

	if err := a.db.Model(&model.Lookup{}).Count(&lookupCount).Error; err != nil {
		return false, fmt.Errorf("failed to count lookups: %v", err)
	}

	// 如果任一表没有数据，则需要初始化
	return apiCount == 0 || lookupCount == 0, nil
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

	err := a.db.Transaction(func(tx *gorm.DB) error {
		for _, it := range lookups {
			if err := tx.FirstOrCreate(it, &model.Lookup{
				GroupValue: it.GroupValue,
				EntryValue: it.EntryValue,
			}).Error; err != nil {
				return fmt.Errorf("failed to create lookup: %v", err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Errorf("lookup initialization failed: %v", err)
	}
}
