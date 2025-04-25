package service

import (
	"admin/internal/dto"
	"admin/internal/model"
	"admin/internal/repository"
	"context"
)

type LookupService interface {
	Create(params dto.LookupCreateRequest) error
	Delete(id uint) error
	Update(id uint, params dto.LookupUpdateRequest) error
	Get(id uint) (*dto.LookupGetIdReqponse, error)
	Status(ctx context.Context, params *dto.LookupStatus) error
	QueryGroup(group_value string, params dto.ListQueryRequest) (dto.GroupQueryResponse, error)
	QueryGroups(params dto.GroupsQueryRequest) (*dto.GroupsQueryResponse, error)
	Sort(params dto.LookupSortRequest) error
}

type lookUpService struct {
	lookupRepo repository.LookupRepository
}

func NewLookupService(lookupRepo repository.LookupRepository) LookupService {
	return &lookUpService{
		lookupRepo: lookupRepo,
	}
}

func (s *lookUpService) Create(params dto.LookupCreateRequest) error {
	return s.lookupRepo.Create(&model.Lookup{
		GroupValue: params.GroupValue,
		EntryLabel: params.EntryLabel,
		EntryValue: params.EntryValue,
		Remark:     params.Remark,
		SortOrder:  params.SortOrder,
		Status:     params.Status,
	})
}

func (s *lookUpService) Delete(id uint) error {
	return s.lookupRepo.Delete(id)
}

func (s *lookUpService) Update(id uint, params dto.LookupUpdateRequest) error {
	return s.lookupRepo.Update(id, &model.Lookup{
		GroupValue: params.GroupValue,
		EntryLabel: params.EntryLabel,
		EntryValue: params.EntryValue,
		Remark:     params.Remark,
		SortOrder:  params.SortOrder,
		Status:     params.Status,
	})
}
func (s *lookUpService) Get(id uint) (*dto.LookupGetIdReqponse, error) {
	it, err := s.lookupRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	var item dto.LookupGetIdReqponse
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
	item.Group = it.GroupValue
	return &item, nil
}

func (s *lookUpService) Status(ctx context.Context, params *dto.LookupStatus) error {
	return s.lookupRepo.Status(ctx, params)
}
func (s *lookUpService) QueryGroup(group_value string, params dto.ListQueryRequest) (dto.GroupQueryResponse, error) {
	return s.lookupRepo.FindLookupGroupsByValue(group_value, params)
}

// 修正方法名，使其与接口一致
func (s *lookUpService) QueryGroups(params dto.GroupsQueryRequest) (*dto.GroupsQueryResponse, error) {
	return s.lookupRepo.FindLookupGroups(&params)
}

// 修正方法名，使其与接口一致
func (s *lookUpService) Sort(params dto.LookupSortRequest) error {
	return s.lookupRepo.Sort(&params)
}
