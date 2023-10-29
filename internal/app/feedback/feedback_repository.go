package feedback

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"

	"gorm.io/gorm"
)

type FeedbackRepositoryImpl struct {
	DB *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) interfaces.FeedbackRepository {
	return &FeedbackRepositoryImpl{DB: db}
}

func (repository *FeedbackRepositoryImpl) Save(feedback *domain.Feedback) (*domain.Feedback, error) {
	result := repository.DB.Create(&feedback)
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}

func (repository *FeedbackRepositoryImpl) Update(feedback *domain.Feedback, id int) (*domain.Feedback, error) {
	result := repository.DB.Where("id = ?", id).Updates(domain.Feedback{FeedbackText: feedback.FeedbackText, FeedbackRating: feedback.FeedbackRating})
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}

func (repository *FeedbackRepositoryImpl) FindById(id int) (*domain.Feedback, error) {
	feedback := domain.Feedback{}

	result := repository.DB.Preload("Customer").First(&feedback, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &feedback, nil
}

func (repository *FeedbackRepositoryImpl) FindAll() ([]domain.Feedback, error) {
	feedback := []domain.Feedback{}

	result := repository.DB.Preload("Customer").Find(&feedback)
	if result.Error != nil {
		return nil, result.Error
	}

	return feedback, nil
}

func (repository *FeedbackRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("feedbacks").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}