package calender

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func PopulateDays(container *fyne.Container, date time.Time, onSelected func(time.Time)) {
	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	startOffset := int(firstDay.Weekday())

	for i := 0; i < startOffset; i++ {
		container.Add(widget.NewLabel(""))
	}

	for day := 1; day <= lastDay.Day(); day++ {
		dayButton := widget.NewButton(
			time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, time.UTC).Format("02"),
			func() {
				onSelected(time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, time.UTC))
			},
		)
		container.Add(dayButton)
	}
}
