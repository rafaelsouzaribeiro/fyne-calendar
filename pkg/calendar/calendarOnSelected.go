package calendar

import (
	"time"
)

func (d *Date) OnSelected(t time.Time) {
	d.Instruction.SetText("Data Selecionada:")
	d.DateChosen.SetText(t.Format("02/01/2006"))
	d.CurrentDate <- t
}
