package handlers

import (
	"askon/support-api/storage"
	"askon/support-api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	articleStorage *storage.ArticleStorage
	ticketStorage  *storage.TicketStorage
}

func NewHandler(articleStorage *storage.ArticleStorage, ticketStorage *storage.TicketStorage) *Handler {
	return &Handler{
		articleStorage: articleStorage,
		ticketStorage:  ticketStorage,
	}
}

func (h *Handler) HandleTicketStates(c *gin.Context) {
	data, err := h.ticketStorage.GetTicketStates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chartData := utils.PrepareChartData(data, "Ticket States", "#4e73df")
	c.JSON(http.StatusOK, chartData)
}

func (h *Handler) HandleArticleTimes(c *gin.Context) {
	fromStr := c.Query("from")
	toStr := c.Query("to")
	if fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "'from' and 'to' parameters are required"})
		return
	}

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'from' date format (use YYYY-MM-DD)"})
		return
	}

	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'to' date format (use YYYY-MM-DD)"})
		return
	}

	data, err := h.articleStorage.GetArticleTime(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chartData := utils.PrepareChartData(data, "Article Creation Times", "#36b9cc")
	c.JSON(http.StatusOK, chartData)
}

func (h *Handler) HandleArticleTypes(c *gin.Context) {
	data, err := h.articleStorage.GetArticleType()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chartData := utils.PrepareChartData(data, "Article Types", "#1cc88a")
	c.JSON(http.StatusOK, chartData)

}
