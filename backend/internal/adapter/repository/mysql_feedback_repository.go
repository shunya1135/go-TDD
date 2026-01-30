package repository

import (
	"abema-discovery/backend/internal/domain/entity"
	"context"
	"database/sql"
	"fmt"
)

type MySQLFeedbackRepository struct {
	db *sql.DB
}

func NewMySQLFeedbackRepository(db *sql.DB) *MySQLFeedbackRepository {
	return &MySQLFeedbackRepository{db: db}
}

func (r *MySQLFeedbackRepository) Save(ctx context.Context, fb *entity.Feedback) error {
	// トランザクション開始
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	defer tx.Rollback()

	//1.feedbacksテーブルに保存
	_, err = tx.ExecContext(ctx, `
	INSET INTO feedbacks (user_id, series_id, feedback_type)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE created_at = NOW()`, fb.UserID, fb.SeriesID, fb.Type)

	if err != nil {
		return err
	}

	//2.feedback_statusテーブルを更新
	column := feedbackTypeToColumn(fb.Type)
	query := fmt.Sprintf(`INSRT INTO feedback_status (series_id, %s, total_count) VALUES (?, 1, 1)
	ON DUPLICATEb KEY UPDATE %s = %s + 1, total_count = total_count +1`, column, column, column)

	_, err = tx.ExecContext(ctx, query, fb.SeriesID)

	if err != nil {
		return err
	}

	// コミット
	return tx.Commit()
}

func feedbackTypeToColumn(t entity.FeedbackType) string {
	switch t {
	case entity.FeedbackHelpful:
		return "helpful_count"

	case entity.FeedbackNotHelpful:
		return "not_helpful_count"

	case entity.FeedbackWatched:
		return "watched_count"

	case entity.FeedbackCompleted:
		return "completed_count"

	case entity.FeedbackDropped:
		return "dropped_count"
	}

	return "helpful_count"

}
