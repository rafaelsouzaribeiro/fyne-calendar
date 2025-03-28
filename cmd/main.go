package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calendário")

	i := widget.NewLabel("Por favor, escolha uma data")
	i.Alignment = fyne.TextAlignCenter
	l := widget.NewLabel("")
	l.Alignment = fyne.TextAlignCenter
	d := &date{instruction: i, dateChosen: l}

	// Define a data inicial do calendário
	startingDate := time.Now()
	calendar := NewTranslatedCalendar(startingDate, d.onSelected)

	c := container.NewVBox(i, l, calendar)

	w.SetContent(c)
	w.ShowAndRun()
}

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label
}

func (d *date) onSelected(t time.Time) {
	// Define o texto em português com o formato desejado
	d.instruction.SetText("Data Selecionada:")
	d.dateChosen.SetText(t.Format("02/01/2006"))
}

// NewTranslatedCalendar cria um calendário traduzido para português
func NewTranslatedCalendar(startingDate time.Time, onSelected func(time.Time)) *fyne.Container {
	currentDate := startingDate
	var calendar *xwidget.Calendar
	// Traduz os nomes dos meses e anos
	monthLabel := widget.NewLabel("")
	updateMonthLabel := func() {
		monthLabel.SetText(translateMonth(currentDate.Month()) + " " + translateYear(currentDate.Year()))
	}

	updateMonthLabel()
	// Criar contêiner para os dias da semana e os dias do mês
	dayLabels := container.NewGridWithColumns(7)
	daysContainer := container.NewGridWithColumns(7)

	// Inicializa os dias da semana
	days := []string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"}
	dayLabels.Objects = nil
	for _, day := range days {
		label := widget.NewLabel(day)
		label.Alignment = fyne.TextAlignCenter
		dayLabels.Add(label)
	}

	// Função para preencher os dias do calendário
	populateDays := func(container *fyne.Container, date time.Time) {
		// Obtém o primeiro dia do mês e o último dia
		firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
		lastDay := firstDay.AddDate(0, 1, -1)
		startOffset := int(firstDay.Weekday()) // Quantos dias "em branco" antes do primeiro dia

		// Adiciona espaços vazios antes do primeiro dia
		for i := 0; i < startOffset; i++ {
			container.Add(widget.NewLabel(""))
		}

		// Adiciona os dias do mês
		for day := 1; day <= lastDay.Day(); day++ {
			dayButton := widget.NewButton(
				time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, time.UTC).Format("02"),
				func() {
					onSelected(time.Date(date.Year(), date.Month(), day, 0, 0, 0, 0, time.UTC))
				},
			)
			container.Add(dayButton)
		}
	}

	// Função para recriar o calendário com a nova data

	updateCalendar := func(date time.Time, calendar *xwidget.Calendar) {
		currentDate = date
		calendar = xwidget.NewCalendar(currentDate, onSelected) // Recria o calendário com a nova data
		updateMonthLabel()                                      // Atualiza o rótulo do mês e ano

		// Limpa e recria os dias do calendário
		daysContainer.Objects = nil
		populateDays(daysContainer, currentDate)
		daysContainer.Refresh()
	}

	// Inicializa os dias do calendário
	populateDays(daysContainer, currentDate)

	// Traduz os botões "Anterior" e "Próximo"
	previousButton := widget.NewButton("Anterior", func() {
		newDate := currentDate.AddDate(0, -1, 0) // Subtrai um mês
		updateCalendar(newDate, calendar)        // Atualiza o calendário
	})
	nextButton := widget.NewButton("Próximo", func() {
		newDate := currentDate.AddDate(0, 1, 0) // Adiciona um mês
		updateCalendar(newDate, calendar)       // Atualiza o calendário
	})

	// Cabeçalho com os botões de navegação e o mês/ano
	header := container.NewVBox(
		container.NewCenter(
			container.NewHBox(previousButton, container.NewCenter(monthLabel), nextButton),
		),
	)

	// Retorna o calendário com o cabeçalho traduzido e os dias abaixo
	return container.NewVBox(header, dayLabels, daysContainer)
}

// translateMonth traduz os nomes dos meses para português
func translateMonth(month time.Month) string {
	months := []string{
		"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
		"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	}
	return months[month-1]
}

// translateYear retorna o ano como string
func translateYear(year int) string {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Format("2006")
}
