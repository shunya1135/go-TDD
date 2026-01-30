package entity

import (
	"errors"
)

type FeedbackType string

const (
	FeedbackHelpful    FeedbackType = "helpful"
	FeedbackNotHelpful FeedbackType = "not_helpful"
	FeedbackWatched    FeedbackType = "watched"
	FeedbackCompleted  FeedbackType = "completed"
	FeedbackDropped    FeedbackType = "dropped"
)

type Feedback struct {
	ID       int64
	UserID   string
	SeriesID string
	Type     FeedbackType
}

func NewFeedback(userID, seriesID string, feedbackType FeedbackType) (*Feedback, error) {
	if userID == "" {
		return nil, errors.New("user_idが空です")
	}
	if seriesID == "" {
		return nil, errors.New("series_idが空です")
	}

	if !isValidFeedbackType(feedbackType) {
		return nil, errors.New("無効なfeedback_typeです")
	}

	return &Feedback{
		UserID:   userID,
		SeriesID: seriesID,
		Type:     feedbackType,
	}, nil
}

func isValidFeedbackType(t FeedbackType) bool {
	switch t {
	case FeedbackHelpful, FeedbackNotHelpful, FeedbackWatched, FeedbackCompleted, FeedbackDropped:
		return true
	}

	return false

}
