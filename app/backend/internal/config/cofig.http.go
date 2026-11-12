package config

import "time"

type HTTPConfig struct {
	Host              string
	Port              string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ShutdownTimeout   time.Duration
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host:              getEnv("HTTP_HOST", "localhost"),
		Port:              getEnv("HTTP_PORT", "8080"),
		ReadTimeout:       getEnvDuration("HTTP_READ_TIMEOUT", 10),
		ReadHeaderTimeout: getEnvDuration("HTTP_READ_HEADER_TIMEOUT", 10),
		WriteTimeout:      getEnvDuration("HTTP_WRITE_TIMEOUT", 10),
		IdleTimeout:       getEnvDuration("HTTP_IDLE_TIMEOUT", 10),
		ShutdownTimeout:   getEnvDuration("HTTP_SHUTDOWN_TIMEOUT", 10),
	}
}
