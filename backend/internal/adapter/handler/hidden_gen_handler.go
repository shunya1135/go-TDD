package handler

import (
	"abema-discovery/backend/internal/usecase"
	"encoding/json"
	"net/http"
)

type HiddenGemHandler struct {
	usecase *usecase.HiddenGemUsecase
}

type HiddenGemResponse struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Genre      string  `json:"genre"`
	Score      float64 `json:"score"`
	WatchCount int     `json:"watch_count"`
	ClickCount int     `json:"click_count"`
	Popularity int     `json:"popularity"`
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

	var response []HiddenGemResponse
	for _, c := range contents {
		score, _ := c.HiddenGemScore()
		response = append(response, HiddenGemResponse{
			ID:         c.ID,
			Title:      c.Title,
			Genre:      c.Genre,
			Score:      score,
			WatchCount: c.WatchCount,
			ClickCount: c.ClickCount,
			Popularity: c.Popularity,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
