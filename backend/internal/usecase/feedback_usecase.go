package usecase

import (
	"abema-discovery/backend/internal/domain/entity"
	"abema-discovery/backend/internal/domain/repository"
	"context"
)

type FeedbackUsecase struct {
	repo repository.FeedbackRepository
}

func NewFeedbackUsecase(repo repository.FeedbackRepository) *FeedbackUsecase {
	return &FeedbackUsecase{repo: repo}
}

func (u *FeedbackUsecase) SubmitFeedback(ctx context.Context, userID, seriesID, feedbackType string) error {

	fb, err := entity.NewFeedback(userID, seriesID, entity.FeedbackType(feedbackType))

	if err != nil {
		return err
	}

	return u.repo.Save(ctx, fb)
}
