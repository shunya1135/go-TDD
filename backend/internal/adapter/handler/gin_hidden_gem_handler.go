package handler

import (
	"abema-discovery/backend/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinHiddenGemHandler struct {
	usecase *usecase.HiddenGemUsecase
}

func NewGinHiddenGemHandler(u *usecase.HiddenGemUsecase) *GinHiddenGemHandler {
	return &GinHiddenGemHandler{usecase: u}
}

func (h *GinHiddenGemHandler) GetHiddenGems(c *gin.Context) {
	// クエリパラメータ取得
	genre := c.Query("genre")
	limitStr := c.Query("limit")

	// Usecaseを呼び出し
	contents, err := h.usecase.GetHiddenGems(c.Request.Context(), genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// limit が指定されていれば件数を制限
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil && limit > 0 && limit < len(contents) {
			contents = contents[:limit]
		}
	}

	c.JSON(http.StatusOK, contents)
}
