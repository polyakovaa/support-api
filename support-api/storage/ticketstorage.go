package storage

import (
	"database/sql"
	"time"
)

type TicketStorage struct {
	db *sql.DB
}

func NewTicketStorage(db *sql.DB) *TicketStorage {
	return &TicketStorage{db: db}
}

func (s *TicketStorage) GetTicketStates(from time.Time, to time.Time) (map[int]int, error) {
	query := ` SELECT ticket_state_id, COUNT(*) as count FROM ticket WHERE create_time BETWEEN ? AND ? GROUP BY ticket_state_id;`

	rows, err := s.db.Query(query, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]int)
	for rows.Next() {
		var stateID, count int
		if err := rows.Scan(&stateID, &count); err != nil {
			return nil, err
		}
		result[stateID] = count
	}

	return result, nil
}
