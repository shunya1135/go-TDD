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

		// Feedback CRUD
		v1 := api.Group("/v1")
		{
			v1.POST("/feedback", feedbackHandler.SubmitFeedback)
			v1.GET("/feedback", feedbackHandler.GetAllFeedbacks)
			v1.GET("/feedback/:id", feedbackHandler.GetFeedbackByID)
			v1.PUT("/feedback/:id", feedbackHandler.UpdateFeedback)
			v1.DELETE("/feedback/:id", feedbackHandler.DeleteFeedback)
		}
	}

	return r
}
