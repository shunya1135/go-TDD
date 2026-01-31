package repository

import (
	"abema-discovery/backend/internal/domain/entity"
	"context"
)

type FeedbackRepository interface {
	// Create
	Save(ctx context.Context, fb *entity.Feedback) error
	// Read
	FindAll(ctx context.Context) ([]*entity.Feedback, error)
	FindByID(ctx context.Context, id int64) (*entity.Feedback, error)
	GetStats(ctx context.Context, seriesID string) (*entity.FeedbackStats, error)
	GetAllStats(ctx context.Context) (map[string]*entity.FeedbackStats, error) // N+1対策
	// Update
	Update(ctx context.Context, fb *entity.Feedback) error
	// Delete
	Delete(ctx context.Context, id int64) error
}
