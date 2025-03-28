package calender

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ChangeDays(dayLabels *fyne.Container, daysContainer *fyne.Container) {

	days := []string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"}
	dayLabels.Objects = nil
	for _, day := range days {
		label := widget.NewLabel(day)
		label.Alignment = fyne.TextAlignCenter
		dayLabels.Add(label)
	}

}
