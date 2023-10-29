package feedback

import (
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type FeedbackHandlerImpl struct {
	FeedbackService interfaces.FeedbackService
}

func NewFeedbackHandler(feedbackService interfaces.FeedbackService) interfaces.FeedbackHandler {
	return &FeedbackHandlerImpl{FeedbackService: feedbackService}
}

func(handler *FeedbackHandlerImpl) CreateFeedbackHandler(ctx echo.Context) error {
	feedbackCreateRequest := web.FeedbackCreateRequest{}
	err := ctx.Bind(&feedbackCreateRequest) 
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.FeedbackService.CreateFeedback(ctx, feedbackCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusCreated(ctx, "Success to create feedback", response)
}


func (handler *FeedbackHandlerImpl) UpdateFeedbackHandler(ctx echo.Context) error {
	feedbackId := ctx.Param("id")
	feedbackIdInt, err := strconv.Atoi(feedbackId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	feedbackUpdateRequest := web.FeedbackUpdateRequest{}
	err = ctx.Bind(&feedbackUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.FeedbackService.UpdateFeedback(ctx, feedbackUpdateRequest, feedbackIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to update feedback", response)
}

func (handler *FeedbackHandlerImpl) GetFeedbackHandler(ctx echo.Context) error {
	feedbackId := ctx.Param("id")
	feedbackIdInt, err := strconv.Atoi(feedbackId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}
	
	response, err := handler.FeedbackService.FindById(ctx, feedbackIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get feedback", response)
}

func (handler *FeedbackHandlerImpl) GetFeedbacksHandler(ctx echo.Context) error {
	response, err := handler.FeedbackService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get feedbacks", response)
}

func (handler *FeedbackHandlerImpl) DeleteFeedbackHandler(ctx echo.Context) error {
	feedbackId := ctx.Param("id")
	feedbackIdInt, err := strconv.Atoi(feedbackId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	err = handler.FeedbackService.DeleteFeedback(ctx, feedbackIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to delete feedback", nil)
}