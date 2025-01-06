package handler

import (
	"admin/internal/dto"
	"admin/internal/service"
	"admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionsHandler struct {
	permissionsService service.PermissionsService
}

func NewPermissionsHandler(permissionsService service.PermissionsService) *PermissionsHandler {
	return &PermissionsHandler{permissionsService: permissionsService}
}

func (s *PermissionsHandler) ParentType(c *gin.Context) {
	var params dto.PermissionsParentRequest
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	data, err := s.permissionsService.FindParentByType(&params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, data)
}
func (s *PermissionsHandler) LIST(c *gin.Context) {
	data, err := s.permissionsService.List()
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, data)
}

func (s *PermissionsHandler) GET(c *gin.Context) {
	var params dto.PermissionsGetIdRequest
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	data, err := s.permissionsService.Get(params.ID)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, data)
}

func (s *PermissionsHandler) POST(c *gin.Context) {
	var params dto.PermissionsCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := s.permissionsService.Create(params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (s *PermissionsHandler) DELETE(c *gin.Context) {
	var params dto.PermissionsDeleteRequest
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := s.permissionsService.Delete(params.ID); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (s *PermissionsHandler) PUT(c *gin.Context) {
	var params dto.PermissionsUpdateRequest
	var idParams dto.PermissionsUpdateIdRequest
	if err := c.ShouldBindUri(&idParams); err != nil {
		response.Fail(c, "id 不能为空")
		return
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := s.permissionsService.Update(idParams.ID, params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (s *PermissionsHandler) Status(c *gin.Context) {
	var params dto.PermissionsStatus
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	err := s.permissionsService.Status(&params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}
