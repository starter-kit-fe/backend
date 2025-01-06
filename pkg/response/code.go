package response

// 错误码常量
const (
	// 系统级别错误码
	SUCCESS           = 200 // 成功
	MOVED_PERMANENTLY = 301 // 重定向
	ERROR             = 500 // 系统错误
	INVALID_PARAMS    = 400 // 参数错误
	UNAUTHORIZED      = 401 // 未授权
	FORBIDDEN         = 403 // 禁止访问
	NOT_FOUND         = 404 // 资源不存在

	// 业务级别错误码 (1000-9999)
	ERROR_USER_NOT_FOUND      = 1001 // 用户不存在
	ERROR_PASSWORD_WRONG      = 1002 // 密码错误
	ERROR_USER_ALREADY_EXISTS = 1003 // 用户已存在
	// ... 更多业务错误码
)

// 错误码对应的消息
var codeMessages = map[int]string{
	SUCCESS:                   "成功",
	ERROR:                     "系统错误",
	INVALID_PARAMS:            "参数错误",
	UNAUTHORIZED:              "未授权",
	FORBIDDEN:                 "禁止访问",
	NOT_FOUND:                 "资源不存在",
	ERROR_USER_NOT_FOUND:      "用户不存在",
	ERROR_PASSWORD_WRONG:      "密码错误",
	ERROR_USER_ALREADY_EXISTS: "用户已存在",
}

// GetMessage 获取错误码对应的消息
func GetMessage(code int) string {
	msg, ok := codeMessages[code]
	if ok {
		return msg
	}
	return "未知错误"
}
