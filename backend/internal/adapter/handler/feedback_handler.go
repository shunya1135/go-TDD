package handler

import (
	"abema-discovery/backend/internal/usecase"
	"encoding/json"
	"net/http"
)

type FeedbackHandler struct {
	usecase *usecase.FeedbackUsecase
}

func NewFeedbackHandler(u *usecase.FeedbackUsecase) *FeedbackHandler {
	return &FeedbackHandler{usecase: u}
}

type FeedbackRequest struct {
	SeriesID     string `json:"series_id"`
	FeedbackType string `json:"feedback_type"`
}

type FeedbackResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (h *FeedbackHandler) SubmitFeedback(w http.ResponseWriter, r *http.Request) {
	// POSTのみ許可
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// リクエストボディをパース
	var req FeedbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 仮のユーザーID（本番では認証から取得）
	userID := "anonymous"

	// Usecaseを呼び出し
	err := h.usecase.SubmitFeedback(r.Context(), userID, req.SeriesID, req.FeedbackType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 成功レスポンス
	response := FeedbackResponse{
		Success: true,
		Message: "フィードバックを記録しました",
	}

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
