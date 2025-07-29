package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/fyne-calendar/pkg/calendar"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calendário")

	i := widget.NewLabel("Por favor, escolha uma data")
	i.Alignment = fyne.TextAlignCenter
	l := widget.NewLabel("")
	l.Alignment = fyne.TextAlignCenter
	d := &calendar.Date{Instruction: i, DateChosen: l}
	startingDate := time.Now()
	specialDays := map[string]bool{
		"2025-07-10": true,
		"2025-07-15": true,
		"2025-07-22": true,
	}
	calendars := calendar.NewTranslatedCalendar(startingDate, specialDays, d.OnSelected)
	c := container.NewVBox(i, l, calendars)
	w.SetContent(c)
	w.ShowAndRun()
}
