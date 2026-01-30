package router

import (
	"abema-discovery/backend/internal/adapter/handler"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(
	hiddenGemHandler *handler.GinHiddenGemHandler,
	feedbackHandler *handler.GinFeedbackHandler,
) *gin.Engine {
	r := gin.Default()

	// API routes
	api := r.Group("/api")
	{
		api.GET("/hidden-gems", hiddenGemHandler.GetHiddenGems)
		api.POST("/v1/feedback", feedbackHandler.SubmitFeedback)
	}

	return r
}
