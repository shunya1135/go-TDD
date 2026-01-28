package usecase

import (
	"abema-discovery/backend/internal/domain/entity"
	"errors"
	"testing"
)

func TestHiddenGemUsecase_GetHiddenGems(t *testing.T) {
	// 正常系：スコア順に並び替えて返す
	t.Run("正常系_スコア順に並び替えて返す", func(t *testing.T) {
		// 準備：モックにデータを設定
		mock := &mockContentRepository{
			contents: []*entity.Content{
				{ID: "1", Title: "作品A", WatchCount: 50, ClickCount: 100, Popularity: 500},
				{ID: "2", Title: "作品B", WatchCount: 80, ClickCount: 100, Popularity: 500},
				{ID: "3", Title: "作品C", WatchCount: 30, ClickCount: 100, Popularity: 500},
			},
		}

		usecase := NewHiddenGemUsecase(mock)

		// 実行
		results, err := usecase.GetHiddenGems("")

		// 検証
		if err != nil {
			t.Errorf("エラーが発生しないはず：%v", err)
		}

		if len(results) != 3 {
			t.Errorf("3件のはず：%d件", len(results))
		}

		// スコア順（B > A > C）
		if results[0].ID != "2" {
			t.Errorf("1番目は作品Bのはず：%s", results[0].ID)
		}

	})

	t.Run("正常系_ジャンル指定で絞り込み", func(t *testing.T) {
		mock := &mockContentRepository{
			contents: []*entity.Content{
				{ID: "1", Genre: "animation", WatchCount: 50, ClickCount: 100, Popularity: 500},
				{ID: "2", Genre: "movie", WatchCount: 80, ClickCount: 100, Popularity: 500},
				{ID: "3", Genre: "animation", WatchCount: 30, ClickCount: 100, Popularity: 500},
			},
		}

		usecase := NewHiddenGemUsecase(mock)

		results, err := usecase.GetHiddenGems("animation")

		if err != nil {
			t.Errorf("エラーが発生しているはず：%v", err)
		}

		// animationは2件
		if len(results) != 2 {
			t.Errorf("2件のはず：%d", len(results))
		}
	})

	t.Run("正常系_0件の場合空リストを返す", func(t *testing.T) {
		mock := &mockContentRepository{
			contents: []*entity.Content{},
		}

		usecase := NewHiddenGemUsecase(mock)

		results, err := usecase.GetHiddenGems("")

		if err != nil {
			t.Errorf("エラーが発生しないはず：%v", err)
		}

		if len(results) != 0 {
			t.Errorf("0件のはず：%d件", len(results))
		}
	})

	t.Run("正常系_スコア計算できない作品は除外", func(t *testing.T) {
		mock := &mockContentRepository{
			contents: []*entity.Content{
				{ID: "1", WatchCount: 50, ClickCount: 100, Popularity: 500}, // 有効
				{ID: "2", WatchCount: 80, ClickCount: 100, Popularity: 50},  // 無効（Popularity < 100）
				{ID: "3", WatchCount: 30, ClickCount: 0, Popularity: 500},   // 無効（ClickCount = 0）
			},
		}

		usecase := NewHiddenGemUsecase(mock)

		results, err := usecase.GetHiddenGems("")

		if err != nil {
			t.Errorf("エラーが発生しないはず：%v", err)
		}

		// 有効な作品は1件だけ
		if len(results) != 1 {
			t.Errorf("1件のはず：%d件", len(results))
		}

		if results[0].ID != "1" {
			t.Errorf("作品1のはず：%s", results[0].ID)
		}
	})

	// 異常系：Repositoryがエラー
	t.Run("異常系_Repositoryエラー", func(t *testing.T) {
		mock := &mockContentRepository{
			err: errors.New("DB接続エラー"),
		}
		usecase := NewHiddenGemUsecase(mock)

		_, err := usecase.GetHiddenGems("")

		if err == nil {
			t.Error("エラーが発生するはず")
		}
	})
}
