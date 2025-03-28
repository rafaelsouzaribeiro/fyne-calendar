package calender

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

type Date struct {
	Instruction *widget.Label
	DateChosen  *widget.Label
}

func (d *Date) OnSelected(t time.Time) {
	d.Instruction.SetText("Data Selecionada:")
	d.DateChosen.SetText(t.Format("02/01/2006"))
}
