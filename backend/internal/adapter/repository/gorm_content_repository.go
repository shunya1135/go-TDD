package repository

import (
	"abema-discovery/backend/internal/domain/entity"

	"gorm.io/gorm"
)

// GORMで使うモデル（DBのテーブル構造）
type SeriesModel struct {
	ID         string `gorm:"column:id;primaryKey"`
	Title      string `gorm:"column:title"`
	GenreID    string `gorm:"column:genre_id"`
	WatchCount int    `gorm:"column:watch_count"`
	ClickCount int    `gorm:"column:click_count"`
	Popularity int    `gorm:"column:popularity"`
}

func (SeriesModel) TableName() string {
	return "series"
}

type GormContentRepository struct {
	db *gorm.DB
}

func NewGormContentRepository(db *gorm.DB) *GormContentRepository {
	return &GormContentRepository{db: db}
}

func (r *GormContentRepository) FindAll() ([]*entity.Content, error) {
	var models []SeriesModel

	// GORMの書き方：SQLを書かない！
	if err := r.db.Where("click_count > 0").Find(&models).Error; err != nil {
		return nil, err
	}

	return toContentEntities(models), nil
}

func (r *GormContentRepository) FindByGenre(genre string) ([]*entity.Content, error) {
	var models []SeriesModel

	if err := r.db.Where("genre_id = ? AND click_count > 0", genre).Find(&models).Error; err != nil {
		return nil, err
	}

	return toContentEntities(models), nil
}

// モデル　→ エンティティ変換
func toContentEntities(models []SeriesModel) []*entity.Content {
	contents := make([]*entity.Content, len(models))
	for i, m := range models {
		contents[i] = &entity.Content{
			ID:         m.ID,
			Title:      m.Title,
			Genre:      m.GenreID,
			WatchCount: m.WatchCount,
			ClickCount: m.ClickCount,
			Popularity: m.Popularity,
		}
	}
	return contents
}
