package calendar

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Date struct {
	Instruction   *widget.Label
	DateChosen    *widget.Label
	StartingDate  time.Time
	SpecialDays   map[string]bool
	CurrentDate   chan time.Time
	DaysContainer *fyne.Container
}

func NewCalendar(date *Date) *Date {
	ch := make(chan time.Time, 1)
	ch <- date.StartingDate

	return &Date{
		Instruction:  date.Instruction,
		DateChosen:   date.DateChosen,
		StartingDate: date.StartingDate,
		SpecialDays:  date.SpecialDays,
		CurrentDate:  ch,
	}

}
