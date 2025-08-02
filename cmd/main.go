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
	calenders := calendar.NewCalendar(&calendar.Date{
		Instruction:  i,
		DateChosen:   l,
		StartingDate: time.Now(),
		SpecialDays:  map[string]bool{"2025-07-10": true, "2025-07-15": true, "2025-07-22": true},
	})
	calendar := calenders.TranslatedCalendar()
	c := container.NewVBox(i, l, calendar)
	w.SetContent(c)
	w.ShowAndRun()
}
