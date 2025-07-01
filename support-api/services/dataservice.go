package services

import (
	"askon/support-api/storage"
)

type DataService struct {
	articleStorage *storage.ArticleStorage
	ticketStorage  *storage.TicketStorage
}

func NewDataService(articleStorage *storage.ArticleStorage, ticketStorage *storage.TicketStorage, batchSize, workers int) *DataService {
	return &DataService{
		articleStorage: articleStorage,
		ticketStorage:  ticketStorage,
	}
}
