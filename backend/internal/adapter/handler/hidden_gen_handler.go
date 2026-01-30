package handler

import (
	"abema-discovery/backend/internal/usecase"
	"encoding/json"
	"net/http"
)

type HiddenGemHandler struct {
	usecase *usecase.HiddenGemUsecase
}

func NewHiddenGemHandler(u *usecase.HiddenGemUsecase) *HiddenGemHandler {
	return &HiddenGemHandler{usecase: u}
}

func (h *HiddenGemHandler) GetHiddenGems(w http.ResponseWriter, r *http.Request) {
	genre := r.URL.Query().Get("genre")

	contents, err := h.usecase.GetHiddenGems(genre)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contents)
}
