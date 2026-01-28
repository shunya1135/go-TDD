package entity

import (
	"errors"
	"math"
)

type Content struct {
	ID         string
	Title      string
	Genre      string
	WatchCount int
	ClickCount int
	Popularity int
}

func (c *Content) EngagementRate() (float64, error) {
	if c.ClickCount == 0 {
		return 0, errors.New("クリック数が0です")
	}
	return float64(c.WatchCount) / float64(c.ClickCount), nil
}

func (c *Content) HiddenGemScore() (float64, error) {
	// 異常系: Popularity < 100
	if c.Popularity < 100 {
		return 0, errors.New("Popularityが100未満です")
	}

	// 異常系: ClickCount = 0
	if c.ClickCount == 0 {
		return 0, errors.New("クリック数が0です")
	}

	// 正常系
	engagementRate := float64(c.WatchCount) / float64(c.ClickCount)
	score := engagementRate * (1 / math.Log10(float64(c.Popularity)+1))

	return score, nil
}
