package dto

import "time"

type ListQueryRequest struct {
	Page      uint       `form:"page" binding:"omitempty,min=1" default:"1"`
	Size      uint       `form:"size" binding:"omitempty,min=1,max=100" default:"10"`
	Sort      string     `form:"sort" binding:"omitempty,oneof=created_at updated_at id" default:"id"`
	Order     string     `form:"order" binding:"omitempty,oneof=asc desc" default:"desc"`
	StartTime *time.Time `form:"startTime" binding:"omitempty" time_format:"2006-01-02T15:04:05Z07:00"`
	EndTime   *time.Time `form:"endTime" binding:"omitempty" time_format:"2006-01-02T15:04:05Z07:00"`
	Status    *uint      `form:"status" binding:"omitempty"`
}
