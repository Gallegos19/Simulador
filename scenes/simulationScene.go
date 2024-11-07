package scenes

import (
	"fmt"
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
        go scene.manageVehicle(vehicle)
        time.Sleep(500 * time.Millisecond)
    }
}

func (scene *SimulationScene) manageVehicle(vehicle *models.Vehicle) {
    for {
        spaceIndex, entered := scene.ParkingLot.Enter(vehicle)
        if entered {
            vehicleView := views.NewVehicleView(vehicle, true)

            scene.updateChan <- func() {
                scene.ParkingLotView.AddVehicle(spaceIndex, vehicleView)
            }

            fmt.Println("Vehículo", vehicle.ID, "ha entrado al estacionamiento.")

            // Simula el tiempo de estacionamiento
            vehicle.Park(2*time.Second, func() {
                scene.ParkingLot.Exit(spaceIndex)
                
                // Envía la actualización de la UI al canal
                scene.updateChan <- func() {
                    scene.ParkingLotView.RemoveVehicle(spaceIndex)
                }

                fmt.Println("Vehículo", vehicle.ID, "ha salido del estacionamiento.")
            })
            return 
        }
        time.Sleep(1 * time.Second)
    }
}


func (scene *SimulationScene) Render() fyne.CanvasObject {
    return scene.Content
}
