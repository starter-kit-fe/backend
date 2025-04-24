package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultPage = 1
	DefaultSize = 10
	MaxSize     = 100
)

// PaginationParams holds the pagination parameters
type PaginationParams struct {
	Page int
	Size int
}

// PaginationMiddleware handles pagination parameters
func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取查询参数
		page := c.Query("page")
		size := c.Query("size")

		// 如果参数为空，设置默认值
		if page == "" {
			c.Request.URL.RawQuery += "&page=" + strconv.Itoa(DefaultPage)
		}
		if size == "" {
			c.Request.URL.RawQuery += "&size=" + strconv.Itoa(DefaultSize)
		}

		// 解析并验证参数
		pageInt, err := strconv.Atoi(c.Query("page"))
		if err != nil || pageInt < 1 {
			c.Request.URL.RawQuery = "page=" + strconv.Itoa(DefaultPage) + "&" + c.Request.URL.RawQuery
		}

		sizeInt, err := strconv.Atoi(c.Query("size"))
		if err != nil || sizeInt < 1 {
			c.Request.URL.RawQuery = "size=" + strconv.Itoa(DefaultSize) + "&" + c.Request.URL.RawQuery
		}

		// 限制最大页大小
		if sizeInt > MaxSize {
			c.Request.URL.RawQuery = "size=" + strconv.Itoa(MaxSize) + "&" + c.Request.URL.RawQuery
		}

		c.Next()
	}
}

// GetPaginationParams retrieves pagination parameters from context
func GetPaginationParams(c *gin.Context) PaginationParams {
	if pagination, exists := c.Get("pagination"); exists {
		return pagination.(PaginationParams)
	}
	return PaginationParams{
		Page: DefaultPage,
		Size: DefaultSize,
	}
}
