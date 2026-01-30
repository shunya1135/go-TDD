package repository

import (
	"abema-discovery/backend/internal/domain/entity"
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// フィードバック用モデル
type FeedbackModel struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement"`
	UserID       string `gorm:"column:user_id"`
	SeriesID     string `gorm:"column:series_id"`
	FeedbackType string `gorm:"column:feedback_type"`
}

func (FeedbackModel) TableName() string {
	return "feedbacks"
}

// フィードバック集計用モデル
type FeedbackStatsModel struct {
	SeriesID        string `gorm:"column:series_id;primaryKey"`
	HelpfulCount    int    `gorm:"column:helpful_count"`
	NotHelpfulCount int    `gorm:"column:not_helpful_count"`
	WatchCount      int    `gorm:"column:watch_count"`
	CompletedCount  int    `gorm:"column:completed_count"`
	DroppedCount    int    `gorm:"column:dropped_count"`
	TotalCount      int    `gorm:"column:total_count"`
}

func (FeedbackStatsModel) TableName() string {
	return "feedback_stats"
}

type GormFeedbackRepository struct {
	db *gorm.DB
}

func NewGormFeedbackRepository(db *gorm.DB) *GormFeedbackRepository {
	return &GormFeedbackRepository{db: db}
}

func (r *GormFeedbackRepository) Save(ctx context.Context, fb *entity.Feedback) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. feedbackテーブルに保存（UPSERT）
		feedback := FeedbackModel{
			UserID:       fb.UserID,
			SeriesID:     fb.SeriesID,
			FeedbackType: string(fb.Type),
		}

		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "series_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"feedback_type"}),
		}).Create(&feedback).Error; err != nil {
			return err
		}

		// 2. feedback_statsを更新
		column := feedbackTypeToColumn(fb.Type)
		stats := FeedbackStatsModel{SeriesID: fb.SeriesID}

		// UPSERT: なければ作成、あれば更新
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "series_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				column:        gorm.Expr(column + " + 1"),
				"total_count": gorm.Expr("total_count + 1"),
			}),
		}).Create(&stats).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *GormFeedbackRepository) GetStats(ctx context.Context, seriesID string) (*entity.FeedbackStats, error) {
	var model FeedbackStatsModel

	err := r.db.Where("series_id = ?", seriesID).First(&model).Error
	if err == gorm.ErrRecordNotFound {
		return &entity.FeedbackStats{SeriesID: seriesID}, nil
	}

	if err != nil {
		return nil, err
	}

	return &entity.FeedbackStats{
		SeriesID:        model.SeriesID,
		HelpfulCount:    model.HelpfulCount,
		NotHelpfulCount: model.NotHelpfulCount,
		WatchedCount:    model.WatchCount,
		CompleteCount:   model.CompletedCount,
		DroppedCount:    model.DroppedCount,
		TotalCount:      model.TotalCount,
	}, nil
}
