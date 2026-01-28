package usecase

import (
	"abema-discovery/backend/internal/domain/entity"
	"abema-discovery/backend/internal/domain/repository"
	"sort"
)

type HiddenGemUsecase struct {
	repo repository.ContentRepository
}

func NewHiddenGemUsecase(repo repository.ContentRepository) *HiddenGemUsecase {
	return &HiddenGemUsecase{repo: repo}
}

func (u *HiddenGemUsecase) GetHiddenGems(genre string) ([]*entity.Content, error) {
	// 作品を取得
	var contents []*entity.Content
	var err error

	if genre == "" {
		contents, err = u.repo.FindAll()
	} else {
		contents, err = u.repo.FindByGenre(genre)
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
