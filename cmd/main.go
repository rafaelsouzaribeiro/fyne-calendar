package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/rafaelsouzaribeiro/internal/calendar"
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
	calendar := calendar.NewTranslatedCalendar(startingDate, d.OnSelected)
	c := container.NewVBox(i, l, calendar)
	w.SetContent(c)
	w.ShowAndRun()
}
