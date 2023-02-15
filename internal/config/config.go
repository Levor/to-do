package config

import (
	"os"
	"strings"
)

type Config struct {
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassword    string
	DbScheme      string
	IsDebugMode   bool
	ServerHost    string
	ServerPort    string
	Cors          *Cors
}

type Cors struct {
	Methods []string
	Origins []string
	Headers []string
}

func Read() *Config {
	cfg := &Config{}
	cfg.DbHost = getEnv("TO_DO_DB_HOST", "localhost")
	cfg.DbPort = getEnv("TO_DO_DB_PORT", "3306")
	cfg.DbUser = getEnv("TO_DO_DB_USER", "root")
	cfg.DbPassword = getEnv("TO_DO_DB_PASS", "secret")
	cfg.DbScheme = getEnv("TO_DO_DB_SCHEME", "todo")
	cfg.ServerHost = getEnv("TO_DO_SERVER_HOST", "localhost")
	cfg.ServerPort = getEnv("TO_DO_SERVER_PORT", "8081")
	cfg.IsDebugMode = true
	cfg.Cors = ReadCorsConfig()

	return cfg
}

func ReadCorsConfig() *Cors {
	cors := Cors{
		Methods: strings.Split(getEnv("TO_DO_CORS_METHODS", "GET,POST,PUT,PATCH,DELETE,OPTIONS"), ","),
		Origins: strings.Split(getEnv("TO_DO_CORS_ORIGINS", "*"), ","),
		Headers: strings.Split(getEnv("TO_DO_CORS_HEADERS", "*"), ","),
	}
	return &cors
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
