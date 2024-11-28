package main

import (
	"log"

	"github.com/weinicm/RighTesting/pkg/db"
)

func main() {
	// 初始化数据库连接
	db, err := db.GetDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	} else {
		log.Println("Database initialized successfully")
	}
	defer db.Close()
}
