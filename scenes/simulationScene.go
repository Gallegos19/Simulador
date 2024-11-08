package scenes

import (
	"time"

	"main/models"
	"main/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SimulationScene struct {
    ParkingLot     *models.ParkingLot
    ParkingLotView *views.ParkingLotView
    EventList      *widget.Label
    Content        *fyne.Container
    updateChan     chan func() 
}


func NewSimulationScene(parkingLot *models.ParkingLot) *SimulationScene {
    parkingLotView := views.NewParkingLotView(parkingLot)

    scene := &SimulationScene{
        ParkingLot:     parkingLot,
        ParkingLotView: parkingLotView,
        EventList:      widget.NewLabel(""),
        updateChan:     make(chan func()),
    }

    scene.Content = container.NewVBox(
        parkingLotView.Render(),
        scene.EventList,
    )

    // Goroutine para manejar actualizaciones de la UI
    go func() {
        for updateFunc := range scene.updateChan {
            updateFunc() 
        }
    }()

 
        time.Sleep(500 * time.Millisecond)
        go scene.runSimulation()
  

    return scene
}

func (scene *SimulationScene) runSimulation() {
    for i := 0; i < 20; i++ {
        vehicle := models.NewVehicle()
        go vehicle.Manage(scene.ParkingLot, scene.updateChan, 
            func(spaceIndex int) { // onEntry
                vehicleView := views.NewVehicleView(vehicle, true)
                scene.updateChan <- func() {
                    scene.ParkingLotView.AddVehicle(spaceIndex, vehicleView)
                }
            }, 
            func(spaceIndex int) { // onExit
                scene.updateChan <- func() {
                    scene.ParkingLotView.RemoveVehicle(spaceIndex)
                }
            })
        time.Sleep(500 * time.Millisecond)
    }
}


func (scene *SimulationScene) Render() fyne.CanvasObject {
    return scene.Content
}
