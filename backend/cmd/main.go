package main

import (
	"abema-discovery/backend/internal/infrastructure/database"
	"log"
)

func main() {
	db, err := database.NewMySQLConnection(database.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "abema",
		Password: "abema123",
		DBName:   "abema_discovery",
	})

	if err != nil {
		log.Fatalf("DB接続エラー：%v", err)
	}

	defer db.Close()

	log.Printf("DB接続成功！")
}
