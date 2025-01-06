package handler

import (
	"admin/internal/dto"
	"admin/internal/service"
	"admin/pkg/response"

	"github.com/gin-gonic/gin"
)

type LookupHandler struct {
	lookUpService service.LookupService
}

func NewLookupHandler(lookUpService service.LookupService) *LookupHandler {
	return &LookupHandler{lookUpService: lookUpService}
}

func (s *LookupHandler) GET(c *gin.Context) {
	var params dto.LookupGetIdRequest
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	data, err := s.lookUpService.Get(params.ID)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, data)
}

func (s *LookupHandler) POST(c *gin.Context) {
	var params dto.LookupCreateRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := s.lookUpService.Create(params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (s *LookupHandler) DELETE(c *gin.Context) {
	var params dto.LookupDeleteRequest
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := s.lookUpService.Delete(params.ID); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (s *LookupHandler) PUT(c *gin.Context) {
	var params dto.LookupUpdateRequest
	var idParams dto.LookupUpdateIdRequest
	if err := c.ShouldBindUri(&idParams); err != nil {
		response.Fail(c, "id 不能为空")
		return
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := s.lookUpService.Update(idParams.ID, params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (s *LookupHandler) Status(c *gin.Context) {
	var params dto.LookupStatus
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	err := s.lookUpService.Status(c.Request.Context(), &params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (s *LookupHandler) Group(c *gin.Context) {
	var params dto.ListQueryRequest
	var GroupValue dto.LookupGroupValue
	if err := c.ShouldBindUri(&GroupValue); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	data, err := s.lookUpService.QueryGroup(GroupValue.GroupValue, params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, data)
}

func (s *LookupHandler) Groups(c *gin.Context) {
	var params dto.GroupsQueryRequest
	if err := c.ShouldBindQuery(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	data, err := s.lookUpService.QueryGroups(params)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, data)

}
