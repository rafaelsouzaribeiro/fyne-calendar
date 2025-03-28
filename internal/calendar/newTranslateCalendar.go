package calendar

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

func NewTranslatedCalendar(startingDate time.Time, onSelected func(time.Time)) *fyne.Container {
	currentDate := startingDate
	var calendar *xwidget.Calendar
	monthLabel := widget.NewLabel("")

	UpdateMonthLabel(monthLabel, currentDate)
	dayLabels := container.NewGridWithColumns(7)
	daysContainer := container.NewGridWithColumns(7)
	ChangeDays(dayLabels, daysContainer)

	updateCalendar := func(date time.Time, calendar *xwidget.Calendar) {
		currentDate = date
		calendar = xwidget.NewCalendar(currentDate, onSelected)
		UpdateMonthLabel(monthLabel, currentDate)
		daysContainer.Objects = nil
		PopulateDays(daysContainer, currentDate, onSelected)
		daysContainer.Refresh()
	}

	PopulateDays(daysContainer, currentDate, onSelected)

	previousButton := widget.NewButton("Anterior", func() {
		newDate := currentDate.AddDate(0, -1, 0)
		updateCalendar(newDate, calendar)
	})
	nextButton := widget.NewButton("Próximo", func() {
		newDate := currentDate.AddDate(0, 1, 0)
		updateCalendar(newDate, calendar)
	})

	header := container.NewVBox(
		container.NewCenter(
			container.NewHBox(previousButton, container.NewCenter(monthLabel), nextButton),
		),
	)

	return container.NewVBox(header, dayLabels, daysContainer)
}
