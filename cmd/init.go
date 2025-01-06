package main

import (
	"context"
	"log"
	"time"
	"admin/config"
	"admin/internal/app"
	"admin/internal/constant"
	"admin/internal/database"
	"admin/pkg/cloudflare"
	"admin/pkg/email"
	"admin/pkg/google"
	"admin/pkg/jwt"
	"admin/pkg/request"
	"admin/pkg/totp"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	params *app.AppMaker
)

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := loadDatabase(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	rdb, err := loadRedis(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	httpClient := createHttpClient()
	turnstile := cloudflare.NewClient(httpClient, "1x0000000000000000000000000000000AA")
	googleService := google.NewGoogleService(httpClient, "1047416391007-femuikjl06pq2tcp40vgs8keg2um3jqu.apps.googleusercontent.com")
	emailClient, err := createEmailService()
	totopClient := totp.NewTOTPGenerator(constant.NAME)
	jwtClient := jwt.NewJWTMaker()
	if err != nil {
		log.Fatal(err.Error())
	}

	params = &app.AppMaker{
		DB:            db,
		RDB:           rdb,
		Request:       httpClient,
		Turnstile:     turnstile,
		GoogleService: googleService,
		EmailClient:   emailClient,
		TotpClient:    totopClient,
		JWT:           jwtClient,
	}
}

func loadDatabase(cfg *config.Config) (*gorm.DB, error) {
	return database.LoadPostgres(cfg.DB.PG)
}

func loadRedis(cfg *config.Config) (*redis.Client, error) {
	return database.LoadRedis(context.Background(), cfg.DB.Redis)
}
func createHttpClient() *request.HttpClient {
	return request.NewHttpClient(
		request.WithTimeout(10*time.Second),
		request.WithUserAgent("admin/1.0"),
		request.WithMaxBodySize(5<<20), // 5MB
		request.WithRetries(3),
	)
}
func createEmailService() (*email.Service, error) {
	return email.NewService(email.Config{
		APIKey:     "re_en3oF5s3_EDmgWsMuLAz9GfLaPnCjQdT4",
		FromEmail:  "noreply@tigerzh.com",
		Domain:     "tigerzh.com",
		TimeFormat: constant.TIME_FORMAT,
	})
}
