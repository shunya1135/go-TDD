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

	feedbackRepo := repository.NewMySQLFeedbackRepository(db)

	// 3.Usecase(Repositoryを渡す)
	uc := usecase.NewHiddenGemUsecase(repo)

	feedbackUC := usecase.NewFeedbackUsecase(feedbackRepo)

	// 4.Handler(Usecaseを渡す)
	h := handler.NewHiddenGemHandler(uc)

	feedbackHandler := handler.NewFeedbackHandler(feedbackUC)

	// 5.Router(Handlerを渡す)
	r := router.NewRouter(h, feedbackHandler)

	// 6.サーバー起動
	log.Println("サーバー起動：http://localhost:8080")
	log.Println("POST /api/v1/feedback でフィードバックを送信")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("サーバーエラー：%v", err)
	}
}
