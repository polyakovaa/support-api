package utils

import (
	"askon/support-api/models"
	"sort"
	"strconv"
)

func PrepareChartData(data interface{}, label string, color string) models.ChartData {
	var labels []string
	var values []int

	switch d := data.(type) {
	case map[int]int:
		// Для числовых ключей сортируем по ключу
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
		// Для строковых ключей (дат) сортируем по дате
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

	return models.ChartData{
		Labels: labels,
		Datasets: []models.Dataset{
			{
				Label:           label,
				Data:            values,
				BackgroundColor: color,
			},
		},
	}
}
