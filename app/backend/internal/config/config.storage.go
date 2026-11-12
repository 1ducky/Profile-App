package config

type StorageConfig struct {
	StorageRoot         string
	StoragePathUpload   string
	StoragePathTemp     string
	StoragePathAvatar   string
	StoragePathDocument string
	StoragePathVideo    string

	// S# Config
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
	Bucket    string
}

func LoadStorageConfig() StorageConfig {
	return StorageConfig{
		StorageRoot:         getEnv("STORAGE_ROOT", "./storage"),
		StoragePathUpload:   getEnv("STORAGE_PATH_UPLOAD", "uploads"),
		StoragePathTemp:     getEnv("STORAGE_PATH_TEMP", "tmp"),
		StoragePathAvatar:   getEnv("STORAGE_PATH_AVATAR", "avatars"),
		StoragePathDocument: getEnv("STORAGE_PATH_DOCUMENT", "documents"),
		StoragePathVideo:    getEnv("STORAGE_PATH_VIDEO", "video"),

		Endpoint:  getEnv("STORAGE_ENDPOINT", "http://localhost:9000"),
		Region:    getEnv("STORAGE_REGION", "us-east-1"),
		AccessKey: getEnv("STORAGE_ACCESS_KEY", "admin"),
		SecretKey: getEnv("STORAGE_SECRET_KEY", "pass1234"),
		Bucket:    getEnv("STORAGE_BUCKET", "profile-app"),
	}
}
