package handler

import (
	"abema-discovery/backend/internal/usecase"
	"net/http"
	"strconv"

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

// SubmitFeedback POST /api/v1/feedback - フィードバックを作成
func (h *GinFeedbackHandler) SubmitFeedback(c *gin.Context) {
	var req GinFeedbackRequst
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userID := req.UserID
	if userID == "" {
		userID = "anonymous"
	}

	err := h.usecase.SubmitFeedback(c.Request.Context(), userID, req.SeriesID, req.FeedbackType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "フィードバックを記録しました",
	})
}

// GetAllFeedbacks GET /api/v1/feedback - 全てのフィードバックを取得
func (h *GinFeedbackHandler) GetAllFeedbacks(c *gin.Context) {
	feedbacks, err := h.usecase.GetAllFeedbacks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}

// GetFeedbackByID GET /api/v1/feedback/:id - IDでフィードバックを取得
func (h *GinFeedbackHandler) GetFeedbackByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	feedback, err := h.usecase.GetFeedbackByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "フィードバックが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, feedback)
}

// UpdateFeedback PUT /api/v1/feedback/:id - フィードバックを更新
func (h *GinFeedbackHandler) UpdateFeedback(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var req struct {
		FeedbackType string `json:"feedback_type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.usecase.UpdateFeedback(c.Request.Context(), id, req.FeedbackType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "フィードバックを更新しました",
	})
}

// DeleteFeedback DELETE /api/v1/feedback/:id - フィードバックを削除
func (h *GinFeedbackHandler) DeleteFeedback(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	err = h.usecase.DeleteFeedback(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "フィードバックを削除しました",
	})
}
