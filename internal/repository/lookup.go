package repository

import (
	"context"
	"fmt"
	"admin/internal/dto"
	"admin/internal/model"
	"admin/pkg/utils"

	"gorm.io/gorm"
)

type LookupRepository interface {
	Create(params *model.Lookup) error
	Delete(id uint) error
	Update(id uint, params *model.Lookup) error
	FindByID(id uint) (*model.Lookup, error)
	Status(ctx context.Context, params *dto.LookupStatus) error
	FindLookupGroupsByValue(group_value string, params dto.ListQueryRequest) (*dto.GroupQueryResponse, error)
	FindLookupGroups(params *dto.GroupsQueryRequest) (*dto.GroupsQueryResponse, error)
}

type lookupRepository struct {
	db       *gorm.DB
	userRepo UserRepository
}

func NewLookupRepository(db *gorm.DB, userRepo UserRepository) LookupRepository {
	return &lookupRepository{
		db:       db,
		userRepo: userRepo,
	}
}

func (s lookupRepository) Create(params *model.Lookup) error {
	return s.db.Create(&params).Error
}

func (s lookupRepository) Delete(id uint) error {
	return s.db.Delete(&model.Lookup{}, id).Error
}

func (s lookupRepository) Update(id uint, params *model.Lookup) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 首先查找要更新的记录
	var lookup model.Lookup
	if err := tx.First(&lookup, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新指定记录
	if err := tx.Model(&lookup).Updates(params).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 如果新的状态是3，则更新同一组内的其他记录状态为1
	if params.Status == 3 {
		if err := tx.Model(&model.Lookup{}).
			Where("group_value = ? AND status = ? AND id != ?", lookup.GroupValue, 3, id).
			Update("status", 1).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
func (s lookupRepository) FindByID(id uint) (*model.Lookup, error) {
	var lookup model.Lookup
	if err := s.db.Find(&lookup, id).Error; err != nil {
		return nil, err
	}
	return &lookup, nil
}

func (s lookupRepository) Status(ctx context.Context, params *dto.LookupStatus) error {
	// 判断状态是不是设为默认
	tx := s.db.WithContext(ctx).Begin()
	var row model.Lookup
	if err := tx.First(&row, params.ID).Error; err != nil {
		return err
	}
	if params.Status == 3 {
		if err := tx.Model(&model.Lookup{}).
			Where("group_value = ? AND status = ? AND id != ?", row.GroupValue, 3, params.ID).
			Update("status", 1).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	result := tx.Model(&model.Lookup{}).
		Where("id = ?", params.ID).
		Update("status", params.Status)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("id错误")
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil

}
func (s lookupRepository) FindLookupGroupsByValue(group_value string, params dto.ListQueryRequest) (*dto.GroupQueryResponse, error) {
	var (
		response dto.GroupQueryResponse
		offset   uint
	)
	query := s.db.Model(&model.Lookup{})
	query = query.Where("group_value = ?", group_value)
	if params.Page != 0 && params.Size != 0 {
		query = utils.BuildBaseQuery(query, params)
		offset = (params.Page - 1) * params.Size
	}

	if err := query.Count(&response.Total).Error; err != nil {
		return nil, err
	}
	var lists []model.Lookup
	if params.Size != 0 {
		query = query.
			Offset(int(offset)).
			Limit(int(params.Size))
	}

	if err := query.
		Find(&lists).Error; err != nil {
		return nil, err
	}
	for _, it := range lists {
		var item dto.GroupQueryResponseItem
		item.ID = it.ID
		item.CreatedAt = it.CreatedAt
		item.UpdatedAt = it.UpdatedAt
		item.IsActive = it.Status == 1 || it.Status == 3
		item.IsDefault = it.Status == 3
		item.Label = it.EntryLabel
		item.Value = it.EntryValue
		item.Sort = it.SortOrder
		item.Status = it.Status
		item.Remark = it.Remark
		if it.CreatedBy != nil {
			if user, err := s.userRepo.FindByID(*it.CreatedBy); err == nil {
				item.Creator = &user.NickName
			}
		}
		if it.UpdatedBy != nil {
			if user, err := s.userRepo.FindByID(*it.UpdatedBy); err == nil {
				item.Creator = &user.NickName
			}
		}
		response.List = append(response.List, item)
	}
	response.Page = params.Page
	return &response, nil
}

func (s lookupRepository) FindLookupGroups(params *dto.GroupsQueryRequest) (*dto.GroupsQueryResponse, error) {
	var response dto.GroupsQueryResponse

	query := s.db.Model(&model.Lookup{})
	// 删除默认查询
	if params.Status != nil && *params.Status != 0 {
		if *params.Status == 1 {
			query = query.Where("status = 1 or status = 3")
		}
		if *params.Status == 2 {
			query = query.Where("status = 2")
		}
	}

	// 查询全局名称
	if params.Name != nil {
		query = query.Where("entry_label LIKE ? OR entry_value LIKE ? OR group_value LIKE ?", "%"+string(*params.Name)+"%", "%"+string(*params.Name)+"%", "%"+string(*params.Name)+"%")
	}
	// 分组
	query = query.Select("group_value, MAX(updated_at) AS updated_at, COUNT(*) AS total").
		Group("group_value").
		Order("MAX(updated_at) DESC")
	offset := (params.Page - 1) * params.Size

	if err := query.Count(&response.Total).Error; err != nil {
		return nil, err
	}

	if err := query.
		Offset(int(offset)).
		Limit(int(params.Size)).
		Find(&response.List).Error; err != nil {
		return nil, err
	}
	response.Page = params.Page
	return &response, nil
}
