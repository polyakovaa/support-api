package handlers

import (
	"askon/support-api/storage"
	"askon/support-api/utils"
	"net/http"

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
	from, to, err := utils.ParseDate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.ticketStorage.GetTicketStates(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chartData := utils.PrepareChartData(data, "Ticket States", "#4e73df")
	c.JSON(http.StatusOK, chartData)
}

func (h *Handler) HandleArticleTimes(c *gin.Context) {
	from, to, err := utils.ParseDate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	from, to, err := utils.ParseDate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.articleStorage.GetArticleType(from, to)

	colors := []string{"#FF6384", "#36A2EB", "#FFCE56", "#4BC0C0", "#9a9fd9", "#ff8cf9", "#ffffa3", "#a3fffa", "#aced91"}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chartData := utils.PrepareChartData(data, "Article Types", colors)
	c.JSON(http.StatusOK, chartData)

}
