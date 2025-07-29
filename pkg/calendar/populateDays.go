package calendar

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PopulateDays(containerObj *fyne.Container, date time.Time, specialDays map[string]bool, onSelected func(time.Time)) {
	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	startOffset := int(firstDay.Weekday())

	for i := 0; i < startOffset; i++ {
		containerObj.Add(widget.NewLabel(""))
	}

	for day := 1; day <= lastDay.Day(); day++ {
		dayButton := widget.NewButton(
			time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, time.UTC).Format("02"),
			func() {
				onSelected(time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, time.UTC))
			},
		)
		dayStr := fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), day)
		var border *canvas.Rectangle
		border = canvas.NewRectangle(color.Transparent)
		if specialDays[dayStr] {
			border = canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 0, A: 255}) // fundo amarelo
			border.StrokeColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}            // borda vermelha
			border.FillColor = color.RGBA{R: 255, G: 255, B: 0, A: 255}            // fundo amarelo
			border.StrokeWidth = 2
		}

		borderedButton := container.NewStack(
			border,
			dayButton,
		)
		containerObj.Add(borderedButton)
	}
}
