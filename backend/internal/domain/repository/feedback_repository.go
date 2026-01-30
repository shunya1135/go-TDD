package repository

import (
	"abema-discovery/backend/internal/domain/entity"
	"context"
)

type FeedbackRepository interface {
	Save(ctx context.Context, fb *entity.Feedback) error
	GetStats(ctx context.Context, seriesID string) (*entity.FeedbackStats, error)
}
