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

func Load() Config {
	if IsDev() {
		return Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				os.Getenv("NEXAMPLE_MYSQL_USER"),
				os.Getenv("NEXAMPLE_MYSQL_PASSWORD"),
				os.Getenv("NEXAMPLE_MYSQL_HOST"),
				os.Getenv("NEXAMPLE_MYSQL_DB"),
			),
			BaseAddress:    "http://localhost:8080",
			SessionAuthKey: "dev-auth-key-32-bytes-long!!!!!!", // 32 bytes
			SessionEncKey:  "dev-encrypt-key-32-bytes-long!!!", // 32 bytes
			CookieName:     "nexample",
		}
	}

	return Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("NEXAMPLE_MYSQL_USER"),
			os.Getenv("NEXAMPLE_MYSQL_PASSWORD"),
			os.Getenv("NEXAMPLE_MYSQL_HOST"),
			os.Getenv("NEXAMPLE_MYSQL_DB"),
		),
		BaseAddress:    os.Getenv("NEXAMPLE_BASE_ADDRESS"),
		SessionAuthKey: os.Getenv("NEXAMPLE_SESSION_AUTH_KEY"),
		SessionEncKey:  os.Getenv("NEXAMPLE_SESSION_ENCRYPT_KEY"),
		CookieName:     "nexample",
	}
}

func IsDev() bool {
	return os.Getenv("NEXAMPLE_ENV") == "development"
}
