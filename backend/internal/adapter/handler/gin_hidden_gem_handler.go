package handler

import (
	"abema-discovery/backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHiddenGemHandler struct {
	usecase *usecase.HiddenGemUsecase
}

func NewGinHiddenGemHandler(u *usecase.HiddenGemUsecase) *GinHiddenGemHandler {
	return &GinHiddenGemHandler{usecase: u}
}

func (h *GinHiddenGemHandler) GetHiddenGems(c *gin.Context) {
	// クエリパラメータ取得（Ginの書き方）
	genre := c.Query("genre")

	// Usecaseを呼び出し（ここは変わらない！）
	contents, err := h.usecase.GetHiddenGems(c.Request.Context(), genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contents)
}
