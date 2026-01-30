package entity

type FeedbackStats struct {
	SeriesID        string
	HelpfulCount    int
	NotHelpfulCount int
	WatchedCount    int
	CompleteCount   int
	DroppedCount    int
	TotalCount      int
}

// ポジティブ率を計算（0.0 〜 1.0）
func (s *FeedbackStats) PositiveRate() float64 {
	if s.TotalCount == 0 {
		return 0.5 // デフォルト
	}
	positive := s.HelpfulCount + s.WatchedCount + s.CompleteCount
	return float64(positive) / float64(s.TotalCount)
}

// FeedbackMultiplier を計算（0.8〜1.2）
func (s *FeedbackStats) Multiplier() float64 {
	if s.TotalCount < 10 {
		return 1.0 //データ不足は補正なし
	}

	// 0% → 0.8, 50% → 1.0, 100% → 1.2
	return 0.8 + (s.PositiveRate() * 0.4)
}
