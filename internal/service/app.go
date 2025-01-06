package service

import (
	"runtime"
	"time"
	"admin/internal/constant"
	"admin/internal/dto"
)

type AppService interface {
	Version() *dto.AppVersionResponse
	Setup() bool
}

type appService struct{}

func NewAppService() AppService {
	return &appService{}
}
func (s *appService) Version() *dto.AppVersionResponse {
	return &dto.AppVersionResponse{
		Now:         time.Now(),
		Version:     constant.VERSION,
		Environment: runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH,
	}
}
func (s *appService) Setup() bool {
	return true
}
