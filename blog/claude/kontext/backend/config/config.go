package config

import "os"

type Config struct {
	Port      string
	JWTSecret string
	DBPath    string
	UploadDir string
}

var AppConfig *Config

func LoadConfig() {
	AppConfig = &Config{
		Port:      getEnv("BLOG_PORT", "8080"),
		JWTSecret: getEnv("BLOG_JWT_SECRET", "blog-secret-key-change-me"),
		DBPath:    getEnv("BLOG_DB_PATH", "./blog.db"),
		UploadDir: getEnv("BLOG_UPLOAD_DIR", "./uploads"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
