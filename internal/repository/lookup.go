package repository

import (
	"admin/internal/dto"
	"admin/internal/model"
	"admin/pkg/utils"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type LookupRepository interface {
	Create(params *model.Lookup) error
	Delete(id uint) error
	Update(id uint, params *model.Lookup) error
	FindByID(id uint) (*model.Lookup, error)
	Status(ctx context.Context, params *dto.LookupStatus) error
	FindLookupGroupsByValue(group_value string, params dto.ListQueryRequest) (dto.GroupQueryResponse, error)
	FindLookupGroups(params *dto.GroupsQueryRequest) (*dto.GroupsQueryResponse, error)
	Sort(params *dto.LookupSortRequest) error
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
func (s lookupRepository) FindLookupGroupsByValue(group_value string, params dto.ListQueryRequest) (dto.GroupQueryResponse, error) {
	var (
		response dto.GroupQueryResponse
	)
	// 初始化 List 为空切片
	response = make([]dto.GroupQueryResponseItem, 0)

	query := s.db.Model(&model.Lookup{})
	var statusMap = map[uint]string{
		0: "status = 0",
		2: "status = 2",
	}
	// 使用处
	if params.Status != nil {
		if condition, ok := statusMap[*params.Status]; ok {
			query = query.Where(condition)
		} else if *params.Status == 1 {
			query = query.Where("status = 1 OR status = 3")
		}
	}
	query = query.Where("group_value = ?", group_value)
	query = query.Order("sort_order asc")
	query = utils.BuildBaseQuery(query, params)
	var lists []model.Lookup

	if err := query.
		Find(&lists).Error; err != nil {
		return response, err
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
		response = append(response, item)
	}
	return response, nil
}

func (s lookupRepository) FindLookupGroups(params *dto.GroupsQueryRequest) (*dto.GroupsQueryResponse, error) {
	var response dto.GroupsQueryResponse
	query := s.db.Model(&model.Lookup{})
	var statusMap = map[uint]string{
		0: "status = 0",
		2: "status = 2",
	}
	// 使用处
	if params.Status != nil {
		if condition, ok := statusMap[*params.Status]; ok {
			query = query.Where(condition)
		} else if *params.Status == 1 {
			query = query.Where("status = 1 OR status = 3")
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

func (s lookupRepository) Sort(params *dto.LookupSortRequest) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for i, it := range params.List {
		if err := tx.Model(&model.Lookup{}).
			Where("id = ?", it.ID).
			Update("sort_order", i+1).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
