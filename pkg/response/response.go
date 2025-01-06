package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code    int          `json:"code"`           // 状态码
	Data    *interface{} `json:"data,omitempty"` // 数据，使用 omitempty 以便在失败时不返回
	Message string       `json:"message"`        // 消息
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Data:    &data,
		Message: GetMessage(SUCCESS),
	})
}

// SuccessWithMessage 自定义消息的成功响应
func SuccessWithMessage(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Data:    &data,
		Message: message,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, message ...string) {
	msg := GetMessage(FORBIDDEN)
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	c.JSON(FORBIDDEN, Response{
		Code:    FORBIDDEN,
		Message: msg,
	})
}

// FailWithCode 自定义状态码的失败响应
func FailWithCode(c *gin.Context, code int, message ...string) {
	msg := GetMessage(code)
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	c.JSON(code, Response{
		Code:    code,
		Message: msg,
	})
}

// CustomResponse 完全自定义的响应
func CustomResponse(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    &data,
		Message: message,
	})
}
