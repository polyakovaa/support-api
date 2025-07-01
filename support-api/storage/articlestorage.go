package storage

import (
	"database/sql"
	"fmt"
	"time"
)

type ArticleStorage struct {
	db *sql.DB
}

func NewArticleStorage(db *sql.DB) *ArticleStorage {
	return &ArticleStorage{db: db}
}

func (a *ArticleStorage) GetArticleTime(from, to time.Time) (map[string]int, error) {
	query := `SELECT 
    	DATE(create_time) as date, 
    	COUNT(*) as messages 
		FROM article 
		WHERE create_time BETWEEN ? AND ?
		GROUP BY date;`

	rows, err := a.db.Query(query, from, to)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var date string
		var count int
		if err := rows.Scan(&date, &count); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		result[date] = count
	}

	return result, nil
}

func (a *ArticleStorage) GetArticleType() (map[int]int, error) {
	query := `SELECT article_type_id, COUNT(*) as count FROM article GROUP BY article_type_id;`

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	result := make(map[int]int)
	for rows.Next() {
		var id, count int
		if err := rows.Scan(&id, &count); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		result[id] = count
	}

	return result, nil
}
