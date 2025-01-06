package dto

import (
	"time"
)

type LookupGetIdRequest struct {
	ID uint `uri:"id" binding:"required"`
}
type LookupGetIdReqponse struct {
	GroupQueryResponseItem
	Group string `json:"group" binding:"required"`
}
type LookupUpdateIdRequest struct {
	LookupGetIdRequest
}
type LookupUpdateRequest struct {
	LookupCreateRequest
}

type LookupDeleteRequest struct {
	LookupGetIdRequest
}

type LookupCreateRequest struct {
	GroupValue string `json:"name" binding:"required"`
	EntryLabel string `json:"label" binding:"required"`
	EntryValue string `json:"value" binding:"required"`
	Remark     string `json:"remark"`
	SortOrder  uint   `json:"sort"  binding:"required"`
	Status     uint   `json:"status" binding:"required"`
}

type LookupStatus struct {
	ID     uint `uri:"id" binding:"required"`
	Status uint `uri:"status" binding:"required"`
}

type LookupGroupValue struct {
	GroupValue string `uri:"group_value" binding:"required"`
}

type GroupQueryResponseItem struct {
	ID        uint      `json:"id"`
	Creator   *string   `json:"creator"`
	Updator   *string   `json:"updator"`
	Value     string    `json:"value"`
	IsDefault bool      `json:"isDefault"`
	IsActive  bool      `json:"isActive"`
	Label     string    `json:"label"`
	Remark    string    `json:"remark"`
	Sort      uint      `json:"sort"`
	Status    uint      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GroupQueryResponse struct {
	List  []GroupQueryResponseItem `json:"list"`
	Total int64                    `json:"total"`
	Page  uint                     `json:"page"`
}

type GroupsQueryRequest struct {
	ListQueryRequest
	Name *string `form:"name" binding:"omitempty,max=100"`
}

type LookupGroupItem struct {
	GroupValue string `json:"value" `
	Total      int64  `json:"total"` // 分组内的记录数

}

type GroupsQueryResponse struct {
	List  []LookupGroupItem `json:"list"`
	Total int64             `json:"total"`
	Page  uint              `json:"page"`
}
