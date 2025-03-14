package handler

import (
	"admin/internal/constant"
	"admin/internal/dto"
	"admin/internal/service"
	"admin/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	tokenExpireDuration = constant.JWT_EXP
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
func (h *UserHandler) GetRouters(c *gin.Context) {
	userIdStr, exists := c.Get("userId")
	if !exists {
		response.FailWithCode(c, response.UNAUTHORIZED)
	}
	userID, ok := userIdStr.(uint)
	if !ok {
		response.FailWithCode(c, response.UNAUTHORIZED)
		return // 确保在类型断言失败时返回
	}
	user, err := h.userService.GetUserRouters(uint(userID))
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, user)
	// c.JSON(http.StatusOK, user)
}
func (h *UserHandler) GoogleSignin(c *gin.Context) {
	var params dto.GoogleSigninRequest
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, err.Error())
		return
	}
	token, err := h.userService.GoogleSignin(c.Request.Context(), params.AccessToken)
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	expiration := time.Now().Add(constant.JWT_EXP)
	c.SetCookie("token", token, int(expiration.Unix()), "/", "", false, false)
	response.Success(c, token)
}

func (h *UserHandler) Signout(c *gin.Context) {
	userIdStr, exists := c.Get("userId")
	if !exists {
		response.FailWithCode(c, response.UNAUTHORIZED)
		return
	}
	userID, ok := userIdStr.(uint)
	if !ok {
		response.FailWithCode(c, response.UNAUTHORIZED)
		return // 确保在类型断言失败时返回
	}
	err := h.userService.Signout(c.Request.Context(), uint(userID))
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "ok")
}

func (h *UserHandler) SendCode(c *gin.Context) {
	var req dto.EmailCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	if err := h.userService.SendVerificationCode(c.Request.Context(), req.Token, req.Email, c.ClientIP()); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "发送成功")
}

func (h *UserHandler) IsEmailExists(c *gin.Context) {
	var params dto.EmailExistsQueryRequest
	if err := c.ShouldBindUri(&params); err != nil {
		response.Fail(c, "请输入正确的邮箱")
		return
	}
	exists := h.userService.IsEmailExists(params.Email)
	response.Success(c, exists)
}

func (h *UserHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "参数错误:"+err.Error())
		return
	}
	token, err := h.userService.Signup(c.Request.Context(), &req, c.ClientIP())
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	expiration := time.Now().Add(constant.JWT_EXP)
	c.SetCookie("token", token, int(expiration.Unix()), "/", "", false, false)
	response.Success(c, token)
}

func (h *UserHandler) Signin(c *gin.Context) {
	var req dto.SigninRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}

	token, err := h.userService.Signin(c.Request.Context(), &req, c.ClientIP())
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	expiration := time.Now().Add(constant.JWT_EXP)
	c.SetCookie("token", token, int(expiration.Unix()), "/", "", false, false)
	response.Success(c, token)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userIdStr, exists := c.Get("userId")
	if !exists {
		response.FailWithCode(c, response.UNAUTHORIZED)
	}
	userID, ok := userIdStr.(uint)
	if !ok {
		response.FailWithCode(c, response.UNAUTHORIZED)
		return // 确保在类型断言失败时返回
	}
	user, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, user)

	// c.JSON(http.StatusOK, user)
}
