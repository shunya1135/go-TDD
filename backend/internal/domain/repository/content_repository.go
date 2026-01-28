package repository

import "abema-discovery/backend/internal/domain/entity"

type ContentRepository interface {
	FindAll() ([]*entity.Content, error)
	FindByGenre(genre string) ([]*entity.Content, error)
}
