package config

import (
	"fmt"
	"os"
)

type Config struct {
	DSN            string
	BaseAddress    string
	SessionAuthKey string
	SessionEncKey  string
	CookieName     string
}

var cfg Config

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("NEXAMPLE_MYSQL_USER", "nexample"),
		getEnv("NEXAMPLE_MYSQL_PASSWORD", ""),
		getEnv("NEXAMPLE_MYSQL_HOST", "127.0.0.1"),
		getEnv("NEXAMPLE_MYSQL_DB", "nexample"),
	)

	if IsDev() {
		cfg = Config{
			DSN:            dsn,
			BaseAddress:    "http://localhost:8080",
			SessionAuthKey: "dev-auth-key-32-bytes-long!!!!!!",
			SessionEncKey:  "dev-encrypt-key-32-bytes-long!!!",
			CookieName:     "nexample",
		}
		return
	}

	cfg = Config{
		DSN:            dsn,
		BaseAddress:    os.Getenv("NEXAMPLE_BASE_ADDRESS"),
		SessionAuthKey: os.Getenv("NEXAMPLE_SESSION_AUTH_KEY"),
		SessionEncKey:  os.Getenv("NEXAMPLE_SESSION_ENCRYPT_KEY"),
		CookieName:     "nexample",
	}
}

func Load() Config {
	return cfg
}

func IsDev() bool {
	return getEnv("NEXAMPLE_ENV", "development") == "development"
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
