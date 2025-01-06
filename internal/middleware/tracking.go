// internal/middleware/tracking.go
package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type TrackingData struct {
	IP        string
	UserAgent string
	UserID    uint
	Path      string
	Method    string
	Operation string
}

// TrackingMiddleware 追踪中间件
func TrackingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取基本信息
		start := time.Now()
		var requestBody bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &requestBody)
		c.Request.Body = io.NopCloser(tee)
		// 2. 继续处理请求
		c.Next()
		time.Since(start).Milliseconds()
		// 3. 如果用户已登录，获取用户ID
		// 4. 异步处理追踪数据
		// 5. 如果不是GET请求，记录操作日志
		if c.Request.Method != "GET" {
			// 记录操作日志
		}
	}
}
