package handler

import (
	"admin/internal/service"
	"admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	appService service.AppService
}

func NewAppHandler(appService service.AppService) *AppHandler {
	return &AppHandler{appService: appService}
}

// Version 获取版本信息
func (h *AppHandler) Version(c *gin.Context) {
	response.Success(c, h.appService.Version())
}

// Version 获取版本信息
func (h *AppHandler) Setup(c *gin.Context) {

	response.Success(c, h.appService.Version())
}
