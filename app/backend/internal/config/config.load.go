package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	StorageConfig StorageConfig
	AuthConfig    AuthConfig
	DBConfig      DBConfig
	HTTPConfig    HTTPConfig
}

func LoadConfig() (AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Print("use callback value")
	}
	StorageConf := LoadStorageConfig()
	authConf := LoadAuthConfig()
	dbConf := LoadMySQLDatabaseConfig()
	HTTPConf := LoadHTTPConfig()
	return AppConfig{StorageConfig: StorageConf, AuthConfig: authConf, DBConfig: dbConf, HTTPConfig: HTTPConf}, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func getEnvDuration(key string, defaultVal int) time.Duration {
	val := getEnv(key, strconv.Itoa(defaultVal))
	n, err := strconv.Atoi(val)
	if err != nil {
		return time.Duration(defaultVal) * time.Second // fall back safely on parse error
	}
	return time.Duration(n) * time.Second
}
