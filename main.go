package main

import (
	"main/models"
	"main/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	mainWindow := myApp.NewWindow("Simulador de Estacionamiento")
	mainWindow.Resize(fyne.NewSize(800, 600))

	parkingLot := models.NewParkingLot(5)
	simulationScene := scenes.NewSimulationScene(parkingLot)

	header := widget.NewLabel("Simulador de Estacionamiento")
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle = fyne.TextStyle{Bold: true}

	statusContainer := container.NewVBox(
		simulationScene.Render(),
	)

	content := container.NewVBox(
		header,
		container.NewHBox(statusContainer),
	)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
