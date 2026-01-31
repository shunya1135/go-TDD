package main

import (
	"abema-discovery/backend/internal/adapter/handler"
	"abema-discovery/backend/internal/adapter/repository"
	"abema-discovery/backend/internal/infrastructure/database"
	"abema-discovery/backend/internal/infrastructure/router"
	"abema-discovery/backend/internal/usecase"
	"log"
)

func main() {
	// 1.DB接続
	// db, err := database.NewMySQLConnection(database.Config{
	// 	Host:     "127.0.0.1",
	// 	Port:     "3306",
	// 	User:     "abema",
	// 	Password: "abema123",
	// 	DBName:   "abema_discovery",
	// })

	// GORM版に変更
	gormDB, err := database.NewGormConnection(database.Config{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "abema",
		Password: "abema123",
		DBName:   "abema_discovery",
	})

	if err != nil {
		log.Fatalf("DB接続エラー：%v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("DB取得エラー：%v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("DB Close エラー：%v", err)
		}
	}()

	// 2.Repository(GORM版)
	contentRepo := repository.NewGormContentRepository(gormDB)
	feedbackRepo := repository.NewGormFeedbackRepository(gormDB)

	// 3.Usecase(Repositoryを渡す)
	contentUC := usecase.NewHiddenGemUsecase(contentRepo, feedbackRepo)

	feedbackUC := usecase.NewFeedbackUsecase(feedbackRepo)

	// 4.Handler(Usecaseを渡す)
	// contentHandler := handler.NewHiddenGemHandler(contentUC)

	// feedbackHandler := handler.NewFeedbackHandler(feedbackUC)

	hiddenGemHandler := handler.NewGinHiddenGemHandler(contentUC)
	feedbackHandler := handler.NewGinFeedbackHandler(feedbackUC)

	// 5.Router(Handlerを渡す)
	// r := router.NewRouter(contentHandler, feedbackHandler)
	r := router.NewGinRouter(hiddenGemHandler, feedbackHandler)

	// 6.サーバー起動
	// log.Println("サーバー起動：http://localhost:8080")
	// log.Println("POST /api/v1/feedback でフィードバックを送信")
	// if err := http.ListenAndServe(":8080", r); err != nil {
	// 	log.Fatalf("サーバーエラー：%v", err)
	// }

	log.Println("サーバー起動： http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("サーバーエラー：%v", err)
	}
}
