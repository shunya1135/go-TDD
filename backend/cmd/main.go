package main

import (
	"abema-discovery/backend/internal/adapter/handler"
	"abema-discovery/backend/internal/adapter/repository"
	"abema-discovery/backend/internal/infrastructure/database"
	"abema-discovery/backend/internal/infrastructure/router"
	"abema-discovery/backend/internal/usecase"
	"log"
	"net/http"
)

func main() {
	// 1.DB接続
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

	// 2.Repository(DBを渡す)
	repo := repository.NewSQLConnectRepository(db)

	// 3.Usecase(Repositoryを渡す)
	uc := usecase.NewHiddenGemUsecase(repo)

	// 4.Handler(Usecaseを渡す)
	h := handler.NewHiddenGemHandler(uc)

	// 5.Router(Handlerを渡す)
	r := router.NewRouter(h)

	// 6.サーバー起動
	log.Println("サーバー起動：http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
