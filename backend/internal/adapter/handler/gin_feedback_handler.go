package handler

import (
	"abema-discovery/backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinFeedbackHandler struct {
	usecase *usecase.FeedbackUsecase
}

func NewGinFeedbackHandler(u *usecase.FeedbackUsecase) *GinFeedbackHandler {
	return &GinFeedbackHandler{usecase: u}
}

type GinFeedbackRequst struct {
	UserID       string `json:"user_id"`
	SeriesID     string `json:"series_id"`
	FeedbackType string `json:"feedback_type"`
}

func (h *GinFeedbackHandler) SubmitFeedback(c *gin.Context) {
	// リクエストボディをパース（Ginの書き方）
	var req GinFeedbackRequst
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})

	}

	// ユーザーID（今回はリクエストから取得）
	userID := req.UserID
	if userID == "" {
		userID = "anonymous"
	}

	// Usecaseを呼び出し（ここは変わらない！）
	err := h.usecase.SubmitFeedback(c.Request.Context(), userID, req.SeriesID, req.FeedbackType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "フィードバックを記録しました",
	})

}
