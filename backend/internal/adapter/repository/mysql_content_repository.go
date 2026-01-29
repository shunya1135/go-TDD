package repository

import (
	"abema-discovery/backend/internal/domain/entity"
	"database/sql"
)

type MySQLConnectRepository struct {
	db *sql.DB
}

func NewSQLConnectRepository(db *sql.DB) *MySQLConnectRepository {
	return &MySQLConnectRepository{db: db}
}

func (r *MySQLConnectRepository) FindAll() ([]*entity.Content, error) {
	query := `SELECT id, title, genre_id, watch_count, click_count, popularity FROM series WHERE click_count > 0`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var contents []*entity.Content

	for rows.Next() {
		c := &entity.Content{}
		err := rows.Scan(&c.ID, &c.Title, &c.Genre, &c.WatchCount, &c.ClickCount, &c.Popularity)

		if err != nil {
			return nil, err
		}

		contents = append(contents, c)
	}
	return contents, nil
}

func (r *MySQLConnectRepository) FindByGenre(genre string) ([]*entity.Content, error) {
	query := `SELECT id, title, genre_id, watch_count, click_count, popularity FROM series WHERE genre_id = ? AND click_count > 0`

	rows, err := r.db.Query(query, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contents []*entity.Content
	for rows.Next() {
		c := &entity.Content{}
		err := rows.Scan(&c.ID, &c.Title, &c.Genre, &c.WatchCount, &c.ClickCount, &c.Popularity)

		if err != nil {
			return nil, err
		}

		contents = append(contents, c)
	}

	return contents, nil
}
