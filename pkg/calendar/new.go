package calendar

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

type Date struct {
	Instruction  *widget.Label
	DateChosen   *widget.Label
	StartingDate time.Time
	SpecialDays  map[string]bool
}

func NewCalendar(date *Date) *Date {
	return &Date{
		Instruction:  date.Instruction,
		DateChosen:   date.DateChosen,
		StartingDate: date.StartingDate,
		SpecialDays:  date.SpecialDays,
	}

}
