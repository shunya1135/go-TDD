package usecase

import (
	"abema-discovery/backend/internal/domain/entity"
	"abema-discovery/backend/internal/domain/repository"
	"context"
	"sort"
)

type HiddenGemUsecase struct {
	contentRepo  repository.ContentRepository
	feedbackRepo repository.FeedbackRepository
}

func NewHiddenGemUsecase(contentRepo repository.ContentRepository,
	feedbackRepo repository.FeedbackRepository) *HiddenGemUsecase {
	return &HiddenGemUsecase{
		contentRepo:  contentRepo,
		feedbackRepo: feedbackRepo,
	}
}

func (u *HiddenGemUsecase) GetHiddenGems(ctx context.Context, genre string) ([]*entity.Content, error) {
	// 作品を取得
	var contents []*entity.Content
	var err error

	if genre == "" {
		contents, err = u.contentRepo.FindAll()
	} else {
		contents, err = u.contentRepo.FindByGenre(genre)
	}

	if err != nil {
		return nil, err
	}

	// スコア計算できる作品だけ抽出
	var validContents []*entity.Content
	for _, c := range contents {
		_, err := c.HiddenGemScore()
		if err == nil {
			validContents = append(validContents, c)
		}
	}

	// スコア順にソート
	sort.Slice(validContents, func(i, j int) bool {
		scoreI, _ := validContents[i].HiddenGemScore()
		scoreJ, _ := validContents[j].HiddenGemScore()
		return scoreI > scoreJ
	})

	return validContents, nil
}

// FinalScore = BaseScore × FeedbackMultiplier
func (u *HiddenGemUsecase) calcFinalScore(ctx context.Context, c *entity.Content) float64 {
	baseScore, _ := c.HiddenGemScore()

	// フィードバック集計を取得
	stats, err := u.feedbackRepo.GetStats(ctx, c.ID)
	if err != nil {
		return baseScore //エラー補正なし
	}

	return baseScore * stats.Multiplier()
}
