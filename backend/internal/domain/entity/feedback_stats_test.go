package entity

import "testing"

func TestFeedbackStats_PositiveRate(t *testing.T) {
	t.Run("正常系_全部ポジティブで100%", func(t *testing.T) {
		stats := FeedbackStats{
			HelpfulCount:  5,
			WatchedCount:  3,
			CompleteCount: 2,
			TotalCount:    10,
		}

		rate := stats.PositiveRate()

		if rate != 1.0 {
			t.Errorf("ポジティブ率は1.0のはず：%v", rate)
		}
	})

	t.Run("異常系_10件未満は補正なし1.0", func(t *testing.T) {
		stats := FeedbackStats{
			HelpfulCount: 5,
			TotalCount:   5,
		}

		m := stats.Multiplier()

		if m != 1.0 {
			t.Errorf("10件未満は1.0のはず：%v", m)
		}
	})
}
