package router

import (
	"abema-discovery/backend/internal/adapter/handler"
	"net/http"
)

func NewRouter(h *handler.HiddenGemHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hidden-gems", h.GetHiddenGems)

	return mux
}
