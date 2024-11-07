package views

import (
	"fmt"
	"main/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ParkingLotView struct {
    ParkingLot   *models.ParkingLot
    Spaces       map[int]*VehicleView 
    Content      *fyne.Container
    InfoLabel    *widget.Label
    FreeSpaceImg *canvas.Image
}

func NewParkingLotView(parkingLot *models.ParkingLot) *ParkingLotView {
    freeSpaceImage := canvas.NewImageFromFile("img/slot2.png")
    freeSpaceImage.FillMode = canvas.ImageFillContain

    view := &ParkingLotView{
        ParkingLot:   parkingLot,
        InfoLabel:    widget.NewLabel("Espacios disponibles: " + strconv.Itoa(parkingLot.Capacity)),
        FreeSpaceImg: freeSpaceImage,
        Spaces:       make(map[int]*VehicleView),
    }

    view.Content = container.NewGridWithColumns(5)
    for i := 0; i < parkingLot.Capacity; i++ {
        space := canvas.NewImageFromFile("img/slot2.png")
        space.FillMode = canvas.ImageFillContain
        space.SetMinSize(fyne.NewSize(80, 50)) 
        view.Content.Add(space)
    }

    return view
}

func (view *ParkingLotView) UpdateParkingLot() {
    for i := 0; i < view.ParkingLot.Capacity; i++ {
        if vehicleView, occupied := view.Spaces[i]; occupied && vehicleView != nil {
            view.Content.Objects[i] = vehicleView.Render()
        } else {
            freeSpaceImage := canvas.NewImageFromFile("img/slot2.png")
            freeSpaceImage.FillMode = canvas.ImageFillContain
            freeSpaceImage.SetMinSize(fyne.NewSize(80, 50)) 
            view.Content.Objects[i] = freeSpaceImage
        }
    }
    availableSpaces := view.ParkingLot.Capacity - len(view.Spaces)
    view.InfoLabel.SetText("Espacios disponibles: " + strconv.Itoa(availableSpaces))
    view.Content.Refresh()
}

func (view *ParkingLotView) RemoveVehicle(spaceIndex int) {
    delete(view.Spaces, spaceIndex)

    freeSpaceImage := canvas.NewImageFromFile("img/slot2.png")
    freeSpaceImage.FillMode = canvas.ImageFillContain
    freeSpaceImage.SetMinSize(fyne.NewSize(80, 50)) 

    if spaceIndex >= 0 && spaceIndex < len(view.Content.Objects) {
        view.Content.Objects[spaceIndex] = freeSpaceImage
    } else {
        fmt.Println("Error: spaceIndex fuera de rango al actualizar la vista")
    }

    availableSpaces := view.ParkingLot.Capacity - len(view.Spaces)
    view.InfoLabel.SetText("Espacios disponibles: " + strconv.Itoa(availableSpaces))

    view.Content.Refresh()
}

func (view *ParkingLotView) AddVehicle(spaceIndex int, vehicleView *VehicleView) {
    view.Spaces[spaceIndex] = vehicleView
    view.UpdateParkingLot()
}




func (view *ParkingLotView) Render() fyne.CanvasObject {
    return container.NewVBox(view.InfoLabel, view.Content)
}
