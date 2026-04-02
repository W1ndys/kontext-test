package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string `mapstructure:"SERVER_PORT" env:"SERVER_PORT" default:":8080"`
	DBPath      string `mapstructure:"DB_PATH" env:"DB_PATH" default:"./data/blog.db"`
	JWTSecret   string `mapstructure:"JWT_SECRET" env:"JWT_SECRET" default:"your-secret-key-change-in-production"`
	JWTExpire   int    `mapstructure:"JWT_EXPIRE" env:"JWT_EXPIRE" default:"72"`
	UploadPath  string `mapstructure:"UPLOAD_PATH" env:"UPLOAD_PATH" default:"./uploads"`
	MaxImageSize int64 `mapstructure:"MAX_IMAGE_SIZE" env:"MAX_IMAGE_SIZE" default:"5242880"`
}

var GlobalConfig *Config

func LoadConfig() *Config {
	if err := os.MkdirAll("./data", 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetDefault("SERVER_PORT", ":8080")
	viper.SetDefault("DB_PATH", "./data/blog.db")
	viper.SetDefault("JWT_SECRET", "your-secret-key-change-in-production")
	viper.SetDefault("JWT_EXPIRE", 72)
	viper.SetDefault("UPLOAD_PATH", "./uploads")
	viper.SetDefault("MAX_IMAGE_SIZE", 5242880)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Warning: Error reading config file: %v", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	if err := os.MkdirAll(config.UploadPath, 0755); err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	GlobalConfig = &config
	fmt.Printf("Config loaded: ServerPort=%s, DBPath=%s, UploadPath=%s\n",
		config.ServerPort, config.DBPath, config.UploadPath)

	return &config
}
