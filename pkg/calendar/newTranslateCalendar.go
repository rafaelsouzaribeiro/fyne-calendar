package calendar

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

func (d *Date) TranslatedCalendar() *fyne.Container {
	currentDate := d.StartingDate
	var calendar *xwidget.Calendar
	monthLabel := widget.NewLabel("")

	d.UpdateMonthLabel(monthLabel, currentDate)
	dayLabels := container.NewGridWithColumns(7)
	d.DaysContainer = container.NewGridWithColumns(7)
	d.ChangeDays(dayLabels, d.DaysContainer)

	updateCalendar := func(date time.Time, calendar *xwidget.Calendar) {
		currentDate = date
		calendar = xwidget.NewCalendar(currentDate, d.OnSelected)
		d.UpdateMonthLabel(monthLabel, currentDate)
		d.DaysContainer.Objects = nil
		d.PopulateDays(d.DaysContainer, currentDate)
		d.DaysContainer.Refresh()
	}

	d.PopulateDays(d.DaysContainer, currentDate)

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

	return container.NewVBox(header, dayLabels, d.DaysContainer)
}
