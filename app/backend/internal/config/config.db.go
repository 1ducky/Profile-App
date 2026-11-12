package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func LoadMySQLDatabaseConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		DbName:   getEnv("DB_NAME", "sgintern-test"),
	}
}
