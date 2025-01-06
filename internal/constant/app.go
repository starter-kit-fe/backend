package constant

import "github.com/gin-gonic/gin"

var (
	VERSION = "N/A"
	MODE    = gin.ReleaseMode
)

const (
	NAME            = "admin"
	DB_TABLE_PREFIX = NAME + "_"
	PORT            = "8000"
	TIME_FORMAT     = "2006-01-02 15:04:05"
	SITE            = "tigerzh.com"
)
