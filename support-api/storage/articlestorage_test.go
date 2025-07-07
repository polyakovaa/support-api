package storage

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestArticleStorage(t *testing.T) {
	t.Run("GetArticleTime success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error creating mock databse: %v", err)
		}
		defer db.Close()

		storage := NewArticleStorage(db)
		from := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		to := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)

		rows := sqlmock.NewRows([]string{"date", "messages"}).
			AddRow("2025-01-01", 10).
			AddRow("2025-01-02", 3)

		mock.ExpectQuery(`SELECT 
    	DATE\(create_time\) as date, 
    	COUNT\(\*\) as messages 
		FROM article 
		WHERE create_time BETWEEN \? AND \?
		GROUP BY date;`).WithArgs(from, to).WillReturnRows(rows)

		result, err := storage.GetArticleTime(from, to)

		assert.NoError(t, err)
		assert.Equal(t, map[string]int{
			"2025-01-01": 10,
			"2025-01-02": 3,
		}, result)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetArticleTime query error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		defer db.Close()

		storage := NewArticleStorage(db)
		from := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
		to := time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC)

		mock.ExpectQuery(`SELECT DATE\(create_time\) as date, COUNT\(\*\) as messages 
		FROM article WHERE create_time BETWEEN \? AND \? 
		GROUP BY date;`).WithArgs(from, to).WillReturnError(errors.New("query failed"))
		_, err = storage.GetArticleTime(from, to)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "query failed")

	})

}
