package feedback

import (
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
)

type FeedbackRoutesImpl struct {
	Echo        *echo.Echo
	FeedbackHandler interfaces.FeedbackHandler
}

func NewFeedbackRoutes(e *echo.Echo, feedbackHandler interfaces.FeedbackHandler) interfaces.FeedbackRoutes {
	return &FeedbackRoutesImpl{
		Echo:        e,
		FeedbackHandler: feedbackHandler,
	}
}

func (routes *FeedbackRoutesImpl) Feedback() {
	feedbacksGroup := routes.Echo.Group("feedbacks")

	feedbacksGroup.POST("", routes.FeedbackHandler.CreateFeedbackHandler)
	feedbacksGroup.GET("", routes.FeedbackHandler.GetFeedbacksHandler)
	feedbacksGroup.GET("/:id", routes.FeedbackHandler.GetFeedbackHandler)
	feedbacksGroup.PUT("/:id", routes.FeedbackHandler.UpdateFeedbackHandler)
	feedbacksGroup.DELETE("/:id", routes.FeedbackHandler.DeleteFeedbackHandler)
}