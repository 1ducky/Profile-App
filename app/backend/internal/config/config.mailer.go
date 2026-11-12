package config

import (
	"strconv"
	"time"
)

type MailerConfig struct {
	Host      string
	Port      int
	User      string
	Password  string
	Identity  string
	Timeout   time.Duration
	MaxConn   int
	QueueSize int
}

func NewMailerConfig() MailerConfig {
	port, err := strconv.Atoi(getEnv("MAILER_PORT", "1025"))
	if err != nil {
		port = 1025
	}
	timeout, err := strconv.Atoi(getEnv("MAILER_TIMEOUT", "10"))
	if err != nil {
		timeout = 10
	}
	queuesize, err := strconv.Atoi(getEnv("MAILER_QUEUE_SIZE", "10"))
	if err != nil {
		queuesize = 10
	}
	return MailerConfig{
		Host:      getEnv("MAILER_HOST", "localhost"),
		Port:      port,
		User:      getEnv("MAILER_USER", "app"),
		Password:  getEnv("MAILER_PASSWORD", "12345678"),
		Identity:  getEnv("MAILER_IDENTITY", "localhost"),
		Timeout:   time.Duration(timeout) * time.Second,
		QueueSize: queuesize,
	}
}
