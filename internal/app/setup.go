package app

import (
	"admin/internal/model"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/gorm"
)

//go:embed lookup.json
var LookUpByte []byte

func (a *App) Setup() {
	if err := a.checkAndMigrateIfNeeded(); err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
}

func (a *App) initApi() error {
	routes := a.router.Routes()
	for _, route := range routes {
		if err := a.db.FirstOrCreate(&model.Api{
			Path:    route.Path,
			Method:  route.Method,
			Handler: route.Handler,
		}, &model.Api{
			Path: route.Path,
		}).Error; err != nil {
			return fmt.Errorf("failed to initialize API route: %v", err)
		}
	}
	return nil
}

func (a *App) checkAndMigrateIfNeeded() error {
	tableExists := a.db.Migrator().HasTable(&model.User{})
	if !tableExists {
		log.Println("Tables don't exist, performing initial migration...")
		if err := a.migrate(); err != nil {
			return fmt.Errorf("migration failed: %v", err)
		}
		if err := a.initApi(); err != nil {
			return err
		}
		if err := a.init(); err != nil {
			return fmt.Errorf("initialization failed: %v", err)
		}
		log.Println("Initial setup completed successfully")
	} else {
		needsInit, err := a.needsInitialization()
		if err != nil {
			return fmt.Errorf("failed to check initialization status: %v", err)
		}
		if needsInit {
			log.Println("Data initialization needed...")
			if err := a.initApi(); err != nil {
				return err
			}
			if err := a.init(); err != nil {
				return fmt.Errorf("data initialization failed: %v", err)
			}
			log.Println("Data initialization completed successfully")
		} else {
			log.Println("Database already initialized, skipping setup")
		}
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

	return apiCount == 0 || lookupCount == 0, nil
}

func (a *App) migrate() error {
	models := []interface{}{
		&model.Api{}, &model.BookArticle{}, &model.Book{}, &model.Categorie{}, &model.Chunk{},
		&model.Dept{}, &model.DeviceIp{}, &model.Device{}, &model.Dialog{}, &model.Lookup{},
		&model.File{}, &model.Ip{}, &model.Log{}, &model.Notice{}, &model.Permissions{},
		&model.PostCategorie{}, &model.PostLike{}, &model.PostView{}, &model.Post{}, &model.RoleMenu{},
		&model.Role{}, &model.UserDevice{}, &model.UserGoogle{}, &model.UserIp{}, &model.UserRole{},
		&model.User{}, &model.UserTotp{},
	}

	if err := a.db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	return nil
}

func (a *App) init() error {
	var lookups []model.Lookup
	if err := json.Unmarshal(LookUpByte, &lookups); err != nil {
		return fmt.Errorf("failed to unmarshal lookup JSON: %v", err)
	}
	return a.db.Transaction(func(tx *gorm.DB) error {
		for _, it := range lookups {
			if err := tx.FirstOrCreate(&it, &model.Lookup{
				GroupValue: it.GroupValue,
				EntryValue: it.EntryValue,
			}).Error; err != nil {
				return fmt.Errorf("failed to create lookup: %v", err)
			}
		}
		return nil
	})
}
