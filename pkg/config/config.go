package config

import (
	"os"
	"strconv"
)

// Config 结构体用于保存数据库配置
type Config struct {
	Database struct {
		Driver   string
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

// LoadConfig 从环境变量加载配置
func LoadConfig() (*Config, error) {
	var config Config

	// 从环境变量中读取配置
	config.Database.Driver = os.Getenv("DATABASE_DRIVER")
	config.Database.Host = os.Getenv("DATABASE_HOST")
	config.Database.Port, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	config.Database.User = os.Getenv("DATABASE_USER")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")
	config.Database.Name = os.Getenv("DATABASE_NAME")

	return &config, nil
}
