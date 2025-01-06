package repository

import (
	"fmt"
	"admin/internal/dto"
	"admin/internal/model"

	"gorm.io/gorm"
)

type PermissionsRepository interface {
	Create(params *model.Permissions) error
	Delete(id uint) error
	Update(id uint, params *model.Permissions) error
	UpdateStatusById(params *dto.PermissionsStatus) error
	FindByID(id uint) (*model.Permissions, error)
	Find() (*[]model.Permissions, error)
	FindParentByType(status uint) (*[]dto.PermissionsParentResponse, error)
}

type permissionsRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionsRepository {
	return &permissionsRepository{
		db: db,
	}
}
func (s *permissionsRepository) FindParentByType(permission_type uint) (*[]dto.PermissionsParentResponse, error) {
	var result []dto.PermissionsParentResponse
	if err := s.db.Model(&model.Permissions{}).
		Select("id, name").
		Where("type = ? AND status = ? ", permission_type, 1).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
func (s *permissionsRepository) UpdateStatusById(params *dto.PermissionsStatus) error {
	if err := s.db.Model(&model.Permissions{}).
		Where("id = ?", params.ID).
		Update("status", params.Status).Error; err != nil {
		return err
	}
	return nil
}

func (s *permissionsRepository) Find() (*[]model.Permissions, error) {
	var response []model.Permissions
	if err := s.db.Find(&response).Error; err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *permissionsRepository) Create(params *model.Permissions) error {
	return s.db.Create(&params).Error
}

func (s *permissionsRepository) Delete(id uint) error {
	var exists bool
	if err := s.db.Model(&model.Permissions{}).
		Select("count(1) > 0").
		Where("parent_id = ?", id).
		Find(&exists).Error; err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("请先删除所有子项")
	}
	return s.db.Delete(&model.Permissions{}, id).Error
}

func (s *permissionsRepository) Update(id uint, params *model.Permissions) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 首先查找要更新的记录
	var permission model.Permissions
	if err := tx.First(&permission, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新指定记录
	if err := tx.Model(&permission).Updates(params).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
func (s *permissionsRepository) FindByID(id uint) (*model.Permissions, error) {
	var lookup model.Permissions
	if err := s.db.Find(&lookup, id).Error; err != nil {
		return nil, err
	}
	return &lookup, nil
}
