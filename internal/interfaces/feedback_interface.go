package interfaces

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"

	"github.com/labstack/echo/v4"
)

type FeedbackRepository interface {
	Save(category *domain.Feedback) (*domain.Feedback, error)
	Update(category *domain.Feedback, id int) (*domain.Feedback, error)
	FindById(id int) (*domain.Feedback, error)
	FindAll() ([]domain.Feedback, error)
	Delete(id int) error
}

type FeedbackService interface {
	CreateFeedback(ctx echo.Context, request web.FeedbackCreateRequest) (*domain.Feedback, error)
	UpdateFeedback(ctx echo.Context, request web.FeedbackUpdateRequest, id int) (*domain.Feedback, error)
	FindById(ctx echo.Context, id int) (*domain.Feedback, error)
	FindAll(ctx echo.Context) ([]domain.Feedback, error)
	DeleteFeedback(ctx echo.Context, id int) error
}

type FeedbackHandler interface {
	CreateFeedbackHandler(ctx echo.Context) error
	UpdateFeedbackHandler(ctx echo.Context) error
	GetFeedbackHandler(ctx echo.Context) error
	GetFeedbacksHandler(ctx echo.Context) error
	DeleteFeedbackHandler(ctx echo.Context) error
}

type FeedbackRoutes interface {
	Feedback()
}