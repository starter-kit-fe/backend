package app

import (
	"admin/pkg/cloudflare"
	"admin/pkg/email"
	"admin/pkg/google"
	"admin/pkg/jwt"
	"admin/pkg/request"
	"admin/pkg/totp"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	db  *gorm.DB
	rdb *redis.Client

	router *gin.Engine

	handlers *Handlers
	services *Services
	repos    *Repositories

	request       *request.HttpClient
	turnstile     *cloudflare.Client
	emailClient   *email.Service
	googleService google.GoogleService
	totpClient    totp.TOTPGenerator
	jwtClient     *jwt.JWTMaker
}
type AppMaker struct {
	DB            *gorm.DB
	RDB           *redis.Client
	Request       *request.HttpClient
	Turnstile     *cloudflare.Client
	EmailClient   *email.Service
	GoogleService google.GoogleService
	TotpClient    totp.TOTPGenerator
	JWT           *jwt.JWTMaker
}

func NewApp(params *AppMaker) *App {
	a := &App{
		db:            params.DB,
		rdb:           params.RDB,
		router:        gin.New(),
		request:       params.Request,
		turnstile:     params.Turnstile,
		emailClient:   params.EmailClient,
		googleService: params.GoogleService,
		totpClient:    params.TotpClient,
		jwtClient:     params.JWT,
	}
	a.router.MaxMultipartMemory = 5 << 20
	// 初始化repositories
	a.initRepositories()
	// 初始化services
	a.initServices()
	// 初始化handlers∏
	a.initHandlers()
	// 初始化中间件
	a.initMiddleware()
	// 初始化路由
	a.initRoutes()
	return a
}
