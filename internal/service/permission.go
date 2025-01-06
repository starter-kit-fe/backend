package service

import (
	"admin/internal/dto"
	"admin/internal/model"
	"admin/internal/repository"
)

type PermissionsService interface {
	Create(params dto.PermissionsCreateRequest) error
	Delete(id uint) error
	Update(id uint, params dto.PermissionsUpdateRequest) error
	Get(id uint) (*dto.PermissionsGetIdReqponse, error)
	List() (*[]dto.PermissionsItemResponse, error)
	Status(params *dto.PermissionsStatus) error
	FindParentByType(params *dto.PermissionsParentRequest) (*[]dto.PermissionsParentResponse, error)
}

type permissionsService struct {
	permissionssRepo repository.PermissionsRepository
	lookupRepo       repository.LookupRepository
}

func NewPermissionsService(permissionssRepo repository.PermissionsRepository, lookupRepo repository.LookupRepository) PermissionsService {
	return &permissionsService{
		permissionssRepo: permissionssRepo,
		lookupRepo:       lookupRepo,
	}
}
func (s *permissionsService) FindParentByType(params *dto.PermissionsParentRequest) (*[]dto.PermissionsParentResponse, error) {
	return s.permissionssRepo.FindParentByType(params.Type)
}
func (s *permissionsService) Status(params *dto.PermissionsStatus) error {
	return s.permissionssRepo.UpdateStatusById(params)
}

func (s *permissionsService) Create(params dto.PermissionsCreateRequest) error {
	return s.permissionssRepo.Create(&model.Permissions{
		Name:     params.Name,
		Type:     params.Type,
		ParentID: params.ParentID,
		Sort:     params.Sort,
		Path:     params.Path,
		IsFrame:  params.IsFrame,
		Status:   params.Status,
		Perms:    params.Perms,
		Icon:     params.Icon,
		Remark:   params.Remark,
	})
}

func (s *permissionsService) Delete(id uint) error {
	return s.permissionssRepo.Delete(id)
}

func (s *permissionsService) Update(id uint, params dto.PermissionsUpdateRequest) error {
	return s.permissionssRepo.Update(id, &model.Permissions{
		Name:     params.Name,
		Type:     params.Type,
		ParentID: params.ParentID,
		Sort:     params.Sort,
		Path:     params.Path,
		IsFrame:  params.IsFrame,
		Status:   params.Status,
		Perms:    params.Perms,
		Icon:     params.Icon,
		Remark:   params.Remark,
	})
}
func (s *permissionsService) Get(id uint) (*dto.PermissionsGetIdReqponse, error) {
	it, err := s.permissionssRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	item := dto.PermissionsGetIdReqponse{
		ID:        it.ID,
		CreatedAt: it.CreatedAt,
		UpdatedAt: it.UpdatedAt,
		Name:      it.Name,
		Path:      it.Path,
		Remark:    it.Remark,
		Sort:      it.Sort,
		Perms:     it.Perms,
		Icon:      it.Icon,
	}
	if data, err := s.lookupRepo.FindByID(it.Type); err == nil {
		item.Type.ID = data.ID
		item.Type.Label = data.EntryLabel
		item.Type.Value = data.EntryValue
	}
	if data, err := s.lookupRepo.FindByID(it.IsFrame); err == nil {
		item.IsFrame.ID = data.ID
		item.IsFrame.Label = data.EntryLabel
		item.IsFrame.Value = data.EntryValue
	}
	if data, err := s.lookupRepo.FindByID(it.Status); err == nil {
		item.Status.ID = data.ID
		item.Status.Label = data.EntryLabel
		item.Status.Value = data.EntryValue
	}
	if data, err := s.permissionssRepo.FindByID(it.ParentID); err == nil {
		item.Parent.ID = data.ID
		item.Parent.Name = data.Name
	}

	return &item, nil
}

func (s *permissionsService) List() (*[]dto.PermissionsItemResponse, error) {
	data, err := s.permissionssRepo.Find()
	if err != nil {
		return nil, err
	}
	var response []dto.PermissionsItemResponse
	for _, it := range *data {
		var item dto.PermissionsItemResponse
		item.ID = it.ID
		item.CreatedAt = it.CreatedAt
		item.UpdatedAt = it.UpdatedAt
		item.Name = it.Name
		if data, err := s.lookupRepo.FindByID(it.Type); err == nil {
			item.Type.ID = data.ID
			item.Type.Label = data.EntryLabel
			item.Type.Value = data.EntryValue
		}
		item.ParentID = it.ParentID
		item.Sort = it.Sort
		item.Path = it.Path
		if data, err := s.lookupRepo.FindByID(it.IsFrame); err == nil {
			item.IsFrame.ID = data.ID
			item.IsFrame.Label = data.EntryLabel
			item.IsFrame.Value = data.EntryValue
		}
		if data, err := s.lookupRepo.FindByID(it.Status); err == nil {
			item.Status.ID = data.ID
			item.Status.Label = data.EntryLabel
			item.Status.Value = data.EntryValue
		}
		item.Perms = it.Perms
		item.Icon = it.Icon
		item.Remark = it.Remark
		response = append(response, item)
		//    it.Type=
	}
	return &response, nil
}
