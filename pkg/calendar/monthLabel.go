package calendar

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

func (d *Date) UpdateMonthLabel(monthLabel *widget.Label, currentDate time.Time) {
	monthLabel.SetText(d.translateMonth(currentDate.Month()) + " " + d.translateYear(currentDate.Year()))
}

func (d *Date) translateMonth(month time.Month) string {
	months := []string{
		"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
		"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	}
	return months[month-1]
}

func (d *Date) translateYear(year int) string {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Format("2006")
}
