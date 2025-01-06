package dto

import (
	"time"
)

type PermissionsParentRequest struct {
	Type uint `uri:"type" binding:"required,min=1"`
}
type PermissionsParentResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PermissionsStatus struct {
	ID     uint `uri:"id" binding:"required"`
	Status uint `uri:"status" binding:"required"`
}
type PermissionsCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Type     uint   `json:"type" binding:"required"`
	ParentID uint   `json:"parentId,omitempty"`
	Sort     uint   `json:"sort" binding:"required"`
	Path     string `json:"path,omitempty" `
	IsFrame  uint   `json:"isFrame,omitempty" `
	Status   uint   `json:"status" binding:"required"`
	Perms    string `json:"perms" binding:"required"`
	Icon     string `json:"icon,omitempty" `
	Remark   string `json:"remark,omitempty"`
}

type PermissionsLookupResponse struct {
	ID    uint   `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

type PermissionsItemResponse struct {
	ID        uint                      `json:"id"`
	CreatedAt time.Time                 `json:"createdAt"`
	UpdatedAt time.Time                 `json:"updatedAt"`
	Name      string                    `json:"name"`
	Type      PermissionsLookupResponse `json:"type"`
	ParentID  uint                      `json:"parentID"`
	Sort      uint                      `json:"sort"`
	Path      string                    `json:"path"`
	IsFrame   PermissionsLookupResponse `json:"isFrame"`
	Status    PermissionsLookupResponse `json:"status"`
	Perms     string                    `json:"perms"`
	Icon      string                    `json:"icon"`
	Remark    string                    `json:"remark"` // 备注

}
type PermissionsDeleteRequest struct {
	PermissionsGetIdRequest
}
type PermissionsUpdateIdRequest struct {
	PermissionsGetIdRequest
}
type PermissionsUpdateRequest struct {
	PermissionsCreateRequest
}
type PermissionsGetIdRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type Parent struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PermissionsGetIdReqponse struct {
	ID        uint                      `json:"id"`
	CreatedAt time.Time                 `json:"createdAt"`
	UpdatedAt time.Time                 `json:"updatedAt"`
	Name      string                    `json:"name"`
	Type      PermissionsLookupResponse `json:"type"`
	Parent    Parent                    `json:"parent"`
	Sort      uint                      `json:"sort"`
	Path      string                    `json:"path"`
	IsFrame   PermissionsLookupResponse `json:"isFrame"`
	Status    PermissionsLookupResponse `json:"status"`
	Perms     string                    `json:"perms"`
	Icon      string                    `json:"icon"`
	Remark    string                    `json:"remark"` // 备注
}
