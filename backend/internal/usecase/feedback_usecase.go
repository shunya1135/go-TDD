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

// GetAllFeedbacks 全てのフィードバックを取得
func (u *FeedbackUsecase) GetAllFeedbacks(ctx context.Context) ([]*entity.Feedback, error) {
	return u.repo.FindAll(ctx)
}

// GetFeedbackByID IDでフィードバックを取得
func (u *FeedbackUsecase) GetFeedbackByID(ctx context.Context, id int64) (*entity.Feedback, error) {
	return u.repo.FindByID(ctx, id)
}

// UpdateFeedback フィードバックを更新
func (u *FeedbackUsecase) UpdateFeedback(ctx context.Context, id int64, feedbackType string) error {
	// 既存のフィードバックを取得
	fb, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// 新しいフィードバックタイプを検証して更新
	newFb, err := entity.NewFeedback(fb.UserID, fb.SeriesID, entity.FeedbackType(feedbackType))
	if err != nil {
		return err
	}
	newFb.ID = id

	return u.repo.Update(ctx, newFb)
}

// DeleteFeedback フィードバックを削除
func (u *FeedbackUsecase) DeleteFeedback(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
