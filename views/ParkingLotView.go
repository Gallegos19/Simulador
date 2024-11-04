package views

import (
    "fyne.io/fyne/v2"
	"strconv" 
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "main/models"
)

type ParkingLotView struct {
    ParkingLot   *models.ParkingLot
    Spaces       []*VehicleView
    Content      *fyne.Container
    InfoLabel    *widget.Label
}

func NewParkingLotView(parkingLot *models.ParkingLot) *ParkingLotView {
    view := &ParkingLotView{
        ParkingLot: parkingLot,
        InfoLabel: widget.NewLabel("Espacios disponibles: 20"),
    }

    // Crea los espacios vacíos en el estacionamiento
    view.Content = container.NewGridWithColumns(5) // Ejemplo: 5 espacios por fila
    for i := 0; i < parkingLot.Capacity; i++ {
        space := widget.NewLabel("Libre")
        view.Content.Add(space)
        view.Spaces = append(view.Spaces, nil) // Espacio vacío al inicio
    }

    return view
}

func (view *ParkingLotView) UpdateParkingLot() {
    for i := 0; i < view.ParkingLot.Capacity; i++ {
        if view.Spaces[i] != nil {
            view.Content.Objects[i] = view.Spaces[i].Render()
        } else {
            view.Content.Objects[i] = widget.NewLabel("Libre")
        }
    }
    view.InfoLabel.SetText("Espacios disponibles: " + strconv.Itoa(view.ParkingLot.Capacity - view.ParkingLot.Occupied))
    view.Content.Refresh()
}

func (view *ParkingLotView) Render() fyne.CanvasObject {
    return container.NewVBox(view.InfoLabel, view.Content)
}
