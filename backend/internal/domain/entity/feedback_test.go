package entity

import (
	"testing"
)

func TestNewFeedback(t *testing.T) {
	t.Run("正常系_有効なfeedback_typeでFeedbackを作成できる", func(t *testing.T) {
		fb, err := NewFeedback("user123", "series456", FeedbackWatched)

		if err != nil {
			t.Errorf("エラーが発生しないはず：%v", err)
		}

		if fb.UserID != "user123" {
			t.Errorf("UserIDが一致しない：%v", fb.UserID)
		}

		if fb.SeriesID != "series456" {
			t.Errorf("SeriesIDが一致しない：%v", fb.SeriesID)
		}

		if fb.Type != FeedbackWatched {
			t.Errorf("Typeが一致しない：%v", fb.Type)
		}
	})

	t.Run("異常系_user_idが空エラー", func(t *testing.T) {
		_, err := NewFeedback("", "series456", FeedbackWatched)

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})

	t.Run("異常系_series_idが空でエラー", func(t *testing.T) {
		_, err := NewFeedback("user123", "", FeedbackWatched)

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}
