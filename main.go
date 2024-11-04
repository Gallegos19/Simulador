package main
//Julio Adrian Gallegos Borraz
import (
	"main/models"
	"main/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
    myApp := app.New()
    myApp.Settings().SetTheme(theme.DarkTheme())

    mainWindow := myApp.NewWindow("Simulador de Estacionamiento")
    mainWindow.Resize(fyne.NewSize(800, 600))

    parkingLot := models.NewParkingLot(20)
    simulationScene := scenes.NewSimulationScene(parkingLot)

    mainWindow.SetContent(container.NewVBox(
        simulationScene.Render(),
    ))   

    mainWindow.ShowAndRun()
}
