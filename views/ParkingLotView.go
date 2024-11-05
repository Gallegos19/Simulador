package views

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "main/models"
    "strconv"
)

type ParkingLotView struct {
    ParkingLot   *models.ParkingLot
    Spaces       []*VehicleView
    Content      *fyne.Container
    InfoLabel    *widget.Label
    FreeSpaceImg *canvas.Image
}

func NewParkingLotView(parkingLot *models.ParkingLot) *ParkingLotView {
    // Cargar la imagen de espacio libre
    freeSpaceImage := canvas.NewImageFromFile("img/slot2.png")
    freeSpaceImage.FillMode = canvas.ImageFillContain

    view := &ParkingLotView{
        ParkingLot:   parkingLot,
        InfoLabel:    widget.NewLabel("Espacios disponibles: 20"),
        FreeSpaceImg: freeSpaceImage,
    }

    // Crea los espacios vacíos en el estacionamiento
    view.Content = container.NewGridWithColumns(5)
    for i := 0; i < parkingLot.Capacity; i++ {
        space := canvas.NewImageFromFile("img/slot2.png")
        space.FillMode = canvas.ImageFillContain
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
            // Usar la imagen de espacio libre cuando no hay un vehículo en el espacio
            freeSpaceImage := canvas.NewImageFromFile("img/slot2.png")
            freeSpaceImage.FillMode = canvas.ImageFillContain
            view.Content.Objects[i] = freeSpaceImage
        }
    }
    view.InfoLabel.SetText("Espacios disponibles: " + strconv.Itoa(view.ParkingLot.Capacity - view.ParkingLot.Occupied))
    view.Content.Refresh()
}

func (view *ParkingLotView) Render() fyne.CanvasObject {
    return container.NewVBox(view.InfoLabel, view.Content)
}
