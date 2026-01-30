package usecase

import (
	"abema-discovery/backend/internal/domain/entity"
	"context"
	"errors"
	"testing"
)

type mockFeedbackRepository struct {
	err error
}

func (m *mockFeedbackRepository) Save(ctx context.Context, fb *entity.Feedback) error {
	return m.err
}

func TestFeedbackUsecase_SubmitFeedback(t *testing.T) {
	t.Run("正常系_フィードバックを保存できる", func(t *testing.T) {
		mock := &mockFeedbackRepository{}
		uc := NewFeedbackUsecase(mock)

		err := uc.SubmitFeedback(context.Background(), "user123", "series456", "watched")

		if err != nil {
			t.Errorf("エラーが発生していないはず：%v", err)
		}
	})

	t.Run("異常系_無効なfeedback_typeでエラー", func(t *testing.T) {
		mock := &mockFeedbackRepository{}
		uc := NewFeedbackUsecase(mock)

		err := uc.SubmitFeedback(context.Background(), "user123", "series456", "invalid")

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})

	t.Run("異常系_Repositoryエラー", func(t *testing.T) {
		mock := &mockFeedbackRepository{err: errors.New("DB error")}
		uc := NewFeedbackUsecase(mock)

		err := uc.SubmitFeedback(context.Background(), "user123", "series456", "watched")

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}
