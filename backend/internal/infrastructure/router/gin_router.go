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
	api := r.Group("/api/v1")
	{
		// Hidden Gems
		api.GET("/hidden-gems", hiddenGemHandler.GetHiddenGems)

		// Feedback CRUD
		api.POST("/feedback", feedbackHandler.SubmitFeedback)
		api.GET("/feedback", feedbackHandler.GetAllFeedbacks)
		api.GET("/feedback/:id", feedbackHandler.GetFeedbackByID)
		api.PUT("/feedback/:id", feedbackHandler.UpdateFeedback)
		api.DELETE("/feedback/:id", feedbackHandler.DeleteFeedback)
	}

	return r
}
