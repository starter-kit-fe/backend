package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DB
}
type DB struct {
	PG    string
	Redis string
}

func LoadConfig() (*Config, error) {
	godotenv.Load()
	DB_URL := os.Getenv("DB_URL")
	REDIS_URL := os.Getenv("REDIS_URL")
	config := Config{
		DB: DB{
			PG:    DB_URL,
			Redis: REDIS_URL,
		},
	}
	if DB_URL == "" {
		return nil, fmt.Errorf("DB_URL is empty, Please set the environment variable like DB_URL='postgresql://postgres:SUq+xmsFg7SwoBwfCYuUFw==@127.0.0.1:5432/admin?sslmode=disable&timezone=Asia/Shanghai'")
	}
	if REDIS_URL == "" {
		return nil, fmt.Errorf("REDIS_URL is empty,  Please set the environment variable like DB_URL='redis://:SUq+xmsFg7SwoBwfCYuUFw==@127.0.0.1:6379/0'")
	}
	return &config, nil
}
