package middleware

import (
	"net/http"
	"sync"
	"time"
	"admin/pkg/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

// UberRateLimiter 为每个IP维护一个限流器
type UberRateLimiter struct {
	limiters map[string]ratelimit.Limiter
	mu       *sync.RWMutex
	rate     int
}

// NewUberRateLimiter 创建一个新的限流器管理器
func NewUberRateLimiter(rate int) *UberRateLimiter {
	return &UberRateLimiter{
		limiters: make(map[string]ratelimit.Limiter),
		mu:       &sync.RWMutex{},
		rate:     rate,
	}
}

// GetLimiter 获取特定IP的限流器
func (u *UberRateLimiter) GetLimiter(ip string) ratelimit.Limiter {
	u.mu.RLock()
	limiter, exists := u.limiters[ip]
	u.mu.RUnlock()

	if !exists {
		u.mu.Lock()
		// 双重检查，避免并发创建
		if limiter, exists = u.limiters[ip]; !exists {
			limiter = ratelimit.New(u.rate) // 每秒允许的请求数
			u.limiters[ip] = limiter
		}
		u.mu.Unlock()
	}

	return limiter
}

// CleanupLimiters 清理长时间未使用的限流器
func (u *UberRateLimiter) CleanupLimiters(maxAge time.Duration) {
	ticker := time.NewTicker(maxAge)
	go func() {
		for range ticker.C {
			u.mu.Lock()
			// 这里可以添加清理逻辑，比如根据最后访问时间
			u.mu.Unlock()
		}
	}()
}

// RateLimitMiddleware 创建一个限流中间件
func RateLimitMiddleware(limiter *UberRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端IP
		clientIP := c.ClientIP()

		// 获取该IP的限流器
		ipLimiter := limiter.GetLimiter(clientIP)

		// Take() 会阻塞直到获取到令牌
		// 如果需要立即返回而不是等待，可以使用下面注释的代码
		ipLimiter.Take()

		/* 非阻塞方式（如果需要可以取消注释使用）
		   now := time.Now()
		   sleepDuration := ipLimiter.Take().Sub(now)
		   if sleepDuration > 0 {
		       c.JSON(http.StatusTooManyRequests, gin.H{
		           "message": "请求太频繁，请稍后再试",
		           "retry_after": sleepDuration.Seconds(),
		       })
		       c.Abort()
		       return
		   }
		*/

		c.Next()
	}
}

// RateLimitConfig 限流配置
type RateLimitConfig struct {
	MaxRequests     int           // 每秒最大请求数
	MaxQueryParams  int           // 最大查询参数数量
	MaxParamLength  int           // 查询参数最大长度
	CleanupInterval time.Duration // 清理间隔
}

// NewRateLimitMiddleware 创建一个包含所有限制的中间件
func NewRateLimitMiddleware(config RateLimitConfig) (gin.HandlerFunc, *UberRateLimiter) {
	limiter := NewUberRateLimiter(config.MaxRequests)

	// 启动清理程序
	if config.CleanupInterval > 0 {
		limiter.CleanupLimiters(config.CleanupInterval)
	}

	return func(c *gin.Context) {
		// 1. 速率限制
		clientIP := c.ClientIP()
		ipLimiter := limiter.GetLimiter(clientIP)
		ipLimiter.Take()

		// 2. 查询参数限制
		if config.MaxQueryParams > 0 {
			queryParams := c.Request.URL.Query()
			if len(queryParams) > config.MaxQueryParams {
				response.FailWithCode(c, http.StatusBadRequest, "查询参数数量超出限制")
				c.Abort()
				return
			}

			if config.MaxParamLength > 0 {
				for _, values := range queryParams {
					for _, v := range values {
						if len(v) > config.MaxParamLength {
							response.FailWithCode(c, http.StatusBadRequest, "查询参数长度超出限制")
							c.Abort()
							return
						}
					}
				}
			}
		}

		c.Next()
	}, limiter
}
