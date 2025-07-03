package utils

import (
	"askon/support-api/models"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrepareChartData(data interface{}, label string, color interface{}) models.ChartData {
	var labels []string
	var values []int

	switch d := data.(type) {
	case map[int]int:
		keys := make([]int, 0, len(d))
		for k := range d {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			labels = append(labels, strconv.Itoa(k))
			values = append(values, d[k])
		}

	case map[string]int:
		dates := make([]string, 0, len(d))
		for date := range d {
			dates = append(dates, date)
		}
		sort.Strings(dates)

		for _, date := range dates {
			labels = append(labels, date)
			values = append(values, d[date])
		}
	}

	var bgColors []string

	switch c := color.(type) {
	case string:
		bgColors = make([]string, len(values))
		for i := range bgColors {
			bgColors[i] = c
		}
	case []string:
		bgColors = c
		if len(bgColors) < len(values) {
			lastColor := bgColors[len(bgColors)-1]
			for len(bgColors) < len(values) {
				bgColors = append(bgColors, lastColor)
			}
		}
	default:
		bgColors = make([]string, len(values))
		for i := range bgColors {
			bgColors[i] = "#36b9cc"
		}
	}

	return models.ChartData{
		Labels: labels,
		Datasets: []models.Dataset{
			{
				Label:           label,
				Data:            values,
				BackgroundColor: bgColors,
			},
		},
	}
}

func ParseDate(c *gin.Context) (from time.Time, to time.Time, err error) {
	fromStr := c.Query("from")
	toStr := c.Query("to")
	if fromStr == "" || toStr == "" {
		return time.Time{}, time.Time{}, fmt.Errorf("'from' and 'to' parameters are required")
	}

	from, err = time.Parse("2006-01-02", fromStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid 'from' date format (use YYYY-MM-DD)")
	}

	to, err = time.Parse("2006-01-02", toStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid 'to' date format (use YYYY-MM-DD)")
	}

	return from, to, nil
}
