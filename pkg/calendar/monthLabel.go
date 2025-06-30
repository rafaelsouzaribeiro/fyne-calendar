package calendar

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

func UpdateMonthLabel(monthLabel *widget.Label, currentDate time.Time) {
	monthLabel.SetText(translateMonth(currentDate.Month()) + " " + translateYear(currentDate.Year()))
}

func translateMonth(month time.Month) string {
	months := []string{
		"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
		"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	}
	return months[month-1]
}

func translateYear(year int) string {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Format("2006")
}
