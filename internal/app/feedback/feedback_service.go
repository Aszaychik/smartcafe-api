package feedback

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/conversion"
	"aszaychik/smartcafe-api/pkg/validation"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type FeedbackServiceImpl struct {
	FeedbackRepository interfaces.FeedbackRepository
	Validate       *validator.Validate
}

func NewFeedbackService(feedbackRepository interfaces.FeedbackRepository, validate *validator.Validate) interfaces.FeedbackService {
	return &FeedbackServiceImpl{
		FeedbackRepository: feedbackRepository,
		Validate:        validate,
	}
}

func (service *FeedbackServiceImpl) CreateFeedback(ctx echo.Context, request web.FeedbackCreateRequest) (*domain.Feedback, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Convert request to domain
	feedback := conversion.FeedbackCreateRequestToFeedbackDomain(request)

	result, err := service.FeedbackRepository.Save(feedback)
	if err != nil {
		return nil, fmt.Errorf("Error when create : %s", err.Error())
	}

	result, _ = service.FeedbackRepository.FindById(int(feedback.ID))

	return result, nil
}

func (service *FeedbackServiceImpl) UpdateFeedback(ctx echo.Context, request web.FeedbackUpdateRequest, id int) (*domain.Feedback, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the feedback exists
	existingFeedback, _ := service.FeedbackRepository.FindById(id)
	if existingFeedback == nil {
		return nil, fmt.Errorf("Feedback not found")
	}

	// Convert request to domain
	feedback := conversion.FeedbackUpdateRequestToFeedbackDomain(request)

	result, err := service.FeedbackRepository.Update(feedback, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating : %s", err.Error())
	}

	result, _ = service.FeedbackRepository.FindById(int(feedback.ID))

	return result, nil
}

func (service *FeedbackServiceImpl) FindById(ctx echo.Context, id int) (*domain.Feedback, error) {
	// Check if the feedback exists
	existingFeedback, _ := service.FeedbackRepository.FindById(id)
	if existingFeedback == nil {
		return nil, fmt.Errorf("Feedback not found")
	}

	return existingFeedback, nil
}

func (service *FeedbackServiceImpl) FindAll(ctx echo.Context) ([]domain.Feedback, error) {
	feedbacks, err := service.FeedbackRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Feedbacks not found")
	}

	return feedbacks, nil
}

func (service *FeedbackServiceImpl) DeleteFeedback(ctx echo.Context, id int) error {
	// Check if the feedback exists
	existingFeedback, _ := service.FeedbackRepository.FindById(id)
	if existingFeedback == nil {
		return fmt.Errorf("Feedback not found")
	}

	err := service.FeedbackRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting : %s", err)
	}

	return nil
}