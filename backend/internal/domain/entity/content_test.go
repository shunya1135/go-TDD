package entity

import (
	"testing"
)

func TestContent_EngagementRate(t *testing.T) {

	// 正常系：視聴数, クリック数100　→　0.5
	t.Run("正常系_エンゲージメント率を計算できる", func(t *testing.T) {
		content := Content{
			WatchCount: 50,
			ClickCount: 100,
		}

		rate, err := content.EngagementRate()

		if err != nil {
			t.Errorf("エラーが発生しないはず：%v", err)
		}

		if rate != 0.5 {
			t.Errorf("期待値 0.5, 実際 %v", rate)
		}

	})

	// 異常系：クリック数0 → エラー
	t.Run("異常系_クリック数0でエラー", func(t *testing.T) {
		content := Content{
			WatchCount: 50,
			ClickCount: 0,
		}

		_, err := content.EngagementRate()

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}

func TestContent_HiddenGemScore(t *testing.T) {
	// 正常系
	t.Run("正常系_スコアを計算できる", func(t *testing.T) {
		content := Content{
			WatchCount: 50,
			ClickCount: 100,
			Popularity: 500,
		}

		score, err := content.HiddenGemScore()

		if err != nil {
			t.Errorf("エラーが発生しないはず：%v", err)
		}

		// スコアが0より大きいことを確認
		if score <= 0 {
			t.Errorf("スコアが0より大きいはず：%v", score)
		}
	})

	// 異常系：Popularity < 100
	t.Run("異常系_Popularityが100未満でエラー", func(t *testing.T) {
		content := Content{
			WatchCount: 50,
			ClickCount: 50,
			Popularity: 0,
		}

		_, err := content.HiddenGemScore()

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})

	// 異常系：ClickCount = 0
	t.Run("異常系_ClickCountが0でエラー", func(t *testing.T) {
		content := Content{
			WatchCount: 50,
			ClickCount: 0,
			Popularity: 50,
		}

		_, err := content.HiddenGemScore()

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}
