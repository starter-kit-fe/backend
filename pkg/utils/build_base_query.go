package utils

import (
	"admin/internal/dto"

	"gorm.io/gorm"
)

func BuildBaseQuery(query *gorm.DB, params dto.ListQueryRequest) *gorm.DB {

	if params.Status != nil && *params.Status != 0 {
		query = query.Where("status = ?", params.Status)
	}
	if params.StartTime != nil {
		query = query.Where("updated_at >= ?", params.StartTime)
	}
	if params.EndTime != nil {
		query = query.Where("updated_at <= ?", params.EndTime)
	}
	if params.Sort != "" && params.Order != "" {
		query = query.Order(params.Sort + " " + params.Order)
	}

	return query
}
