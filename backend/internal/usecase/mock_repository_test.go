package usecase

import (
	"abema-discovery/backend/internal/domain/entity"
)

// モック：テスト用の偽Repository
type mockContentRepository struct {
	contents []*entity.Content
	err      error
}

func (m *mockContentRepository) FindAll() ([]*entity.Content, error) {
	return m.contents, m.err
}

func (m *mockContentRepository) FindByGenre(genre string) ([]*entity.Content, error) {
	if m.err != nil {
		return nil, m.err
	}

	var result []*entity.Content
	for _, c := range m.contents {
		if c.Genre == genre {
			result = append(result, c)
		}
	}

	return result, nil
}
