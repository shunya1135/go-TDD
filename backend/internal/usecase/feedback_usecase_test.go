package usecase

import (
	"abema-discovery/backend/internal/domain/entity"
	"context"
	"errors"
	"testing"
)

type mockFeedbackRepository struct {
	feedbacks []*entity.Feedback
	err       error
	deleteErr error
}

func (m *mockFeedbackRepository) Save(ctx context.Context, fb *entity.Feedback) error {
	return m.err
}

func (m *mockFeedbackRepository) GetStats(ctx context.Context, seriesID string) (*entity.FeedbackStats, error) {
	return &entity.FeedbackStats{}, nil
}

func (m *mockFeedbackRepository) FindAll(ctx context.Context) ([]*entity.Feedback, error) {
	return m.feedbacks, m.err
}

func (m *mockFeedbackRepository) FindByID(ctx context.Context, id int64) (*entity.Feedback, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, fb := range m.feedbacks {
		if fb.ID == id {
			return fb, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockFeedbackRepository) Update(ctx context.Context, fb *entity.Feedback) error {
	return m.err
}

func (m *mockFeedbackRepository) Delete(ctx context.Context, id int64) error {
	return m.deleteErr
}

func (m *mockFeedbackRepository) GetAllStats(ctx context.Context) (map[string]*entity.FeedbackStats, error) {
	return make(map[string]*entity.FeedbackStats), m.err
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

func TestFeedbackUsecase_GetAllFeedbacks(t *testing.T) {
	t.Run("正常系_全てのフィードバックを取得できる", func(t *testing.T) {
		mock := &mockFeedbackRepository{
			feedbacks: []*entity.Feedback{
				{ID: 1, UserID: "user1", SeriesID: "series1", Type: entity.FeedbackHelpful},
				{ID: 2, UserID: "user2", SeriesID: "series2", Type: entity.FeedbackWatched},
			},
		}
		uc := NewFeedbackUsecase(mock)

		feedbacks, err := uc.GetAllFeedbacks(context.Background())

		if err != nil {
			t.Errorf("エラーが発生しないはず: %v", err)
		}
		if len(feedbacks) != 2 {
			t.Errorf("2件取得されるはず: got %d", len(feedbacks))
		}
	})

	t.Run("正常系_0件の場合空リストを返す", func(t *testing.T) {
		mock := &mockFeedbackRepository{feedbacks: []*entity.Feedback{}}
		uc := NewFeedbackUsecase(mock)

		feedbacks, err := uc.GetAllFeedbacks(context.Background())

		if err != nil {
			t.Errorf("エラーが発生しないはず: %v", err)
		}
		if len(feedbacks) != 0 {
			t.Errorf("0件のはず: got %d", len(feedbacks))
		}
	})

	t.Run("異常系_Repositoryエラー", func(t *testing.T) {
		mock := &mockFeedbackRepository{err: errors.New("DB error")}
		uc := NewFeedbackUsecase(mock)

		_, err := uc.GetAllFeedbacks(context.Background())

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}

func TestFeedbackUsecase_GetFeedbackByID(t *testing.T) {
	t.Run("正常系_IDでフィードバックを取得できる", func(t *testing.T) {
		mock := &mockFeedbackRepository{
			feedbacks: []*entity.Feedback{
				{ID: 1, UserID: "user1", SeriesID: "series1", Type: entity.FeedbackHelpful},
			},
		}
		uc := NewFeedbackUsecase(mock)

		feedback, err := uc.GetFeedbackByID(context.Background(), 1)

		if err != nil {
			t.Errorf("エラーが発生しないはず: %v", err)
		}
		if feedback.ID != 1 {
			t.Errorf("ID=1のはず: got %d", feedback.ID)
		}
	})

	t.Run("異常系_存在しないID", func(t *testing.T) {
		mock := &mockFeedbackRepository{feedbacks: []*entity.Feedback{}}
		uc := NewFeedbackUsecase(mock)

		_, err := uc.GetFeedbackByID(context.Background(), 999)

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}

func TestFeedbackUsecase_UpdateFeedback(t *testing.T) {
	t.Run("正常系_フィードバックを更新できる", func(t *testing.T) {
		mock := &mockFeedbackRepository{
			feedbacks: []*entity.Feedback{
				{ID: 1, UserID: "user1", SeriesID: "series1", Type: entity.FeedbackHelpful},
			},
		}
		uc := NewFeedbackUsecase(mock)

		err := uc.UpdateFeedback(context.Background(), 1, "watched")

		if err != nil {
			t.Errorf("エラーが発生しないはず: %v", err)
		}
	})

	t.Run("異常系_無効なfeedback_type", func(t *testing.T) {
		mock := &mockFeedbackRepository{
			feedbacks: []*entity.Feedback{
				{ID: 1, UserID: "user1", SeriesID: "series1", Type: entity.FeedbackHelpful},
			},
		}
		uc := NewFeedbackUsecase(mock)

		err := uc.UpdateFeedback(context.Background(), 1, "invalid_type")

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})

	t.Run("異常系_存在しないID", func(t *testing.T) {
		mock := &mockFeedbackRepository{feedbacks: []*entity.Feedback{}}
		uc := NewFeedbackUsecase(mock)

		err := uc.UpdateFeedback(context.Background(), 999, "watched")

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}

func TestFeedbackUsecase_DeleteFeedback(t *testing.T) {
	t.Run("正常系_フィードバックを削除できる", func(t *testing.T) {
		mock := &mockFeedbackRepository{}
		uc := NewFeedbackUsecase(mock)

		err := uc.DeleteFeedback(context.Background(), 1)

		if err != nil {
			t.Errorf("エラーが発生しないはず: %v", err)
		}
	})

	t.Run("異常系_Repositoryエラー", func(t *testing.T) {
		mock := &mockFeedbackRepository{deleteErr: errors.New("DB error")}
		uc := NewFeedbackUsecase(mock)

		err := uc.DeleteFeedback(context.Background(), 1)

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}
