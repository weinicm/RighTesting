package db

import (
	"fmt"
	"sync"

	db "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"github.com/weinicm/RighTesting/pkg/config"
)

var (
	once     sync.Once
	instance db.Session
	dbConfig *config.Config
)

func init() {
	// Load configuration during package initialization
	var err error
	dbConfig, err = config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
}

// GetDB 返回单例数据库连接
func GetDB() (db.Session, error) {
	var err error
	once.Do(func() {
		if dbConfig == nil {
			err = fmt.Errorf("config is not set")
			return
		}

		settings := postgresql.ConnectionURL{
			Database: dbConfig.Database.Name,
			Host:     dbConfig.Database.Host,
			Password: dbConfig.Database.Password,
			User:     dbConfig.Database.User,
		}

		instance, err = postgresql.Open(settings)
		if err != nil {
			err = fmt.Errorf("db.Open(): %w", err)
		}
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}

// Close 关闭数据库连接
func Close() error {
	if instance != nil {
		return instance.Close()
	}
	return nil
}
