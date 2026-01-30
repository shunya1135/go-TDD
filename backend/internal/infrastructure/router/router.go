package router

import (
	"abema-discovery/backend/internal/adapter/handler"
	"net/http"
)

func NewRouter(h *handler.HiddenGemHandler, feedbackHanlder *handler.FeedbackHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// 隠れ名作API
	mux.HandleFunc("/api/hidden-gems", h.GetHiddenGems)

	// フィードバックAPI
	mux.HandleFunc("/api/v1/feedback", feedbackHanlder.SubmitFeedback)

	return mux
}
