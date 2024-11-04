package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"main/models"
)

type VehicleView struct {
	Vehicle *models.Vehicle
	Image   *canvas.Image
}

// Hide implements fyne.CanvasObject.
func (v *VehicleView) Hide() {
	panic("unimplemented")
}

// MinSize implements fyne.CanvasObject.
func (v *VehicleView) MinSize() fyne.Size {
	panic("unimplemented")
}

// Move implements fyne.CanvasObject.
func (v *VehicleView) Move(fyne.Position) {
	panic("unimplemented")
}

// Position implements fyne.CanvasObject.
func (v *VehicleView) Position() fyne.Position {
	panic("unimplemented")
}

// Refresh implements fyne.CanvasObject.
func (v *VehicleView) Refresh() {
	panic("unimplemented")
}

// Resize implements fyne.CanvasObject.
func (v *VehicleView) Resize(fyne.Size) {
	panic("unimplemented")
}

// Show implements fyne.CanvasObject.
func (v *VehicleView) Show() {
	panic("unimplemented")
}

// Size implements fyne.CanvasObject.
func (v *VehicleView) Size() fyne.Size {
	panic("unimplemented")
}

// Visible implements fyne.CanvasObject.
func (v *VehicleView) Visible() bool {
	panic("unimplemented")
}

func NewVehicleView(vehicle *models.Vehicle, facingRight bool) *VehicleView {
	imagePath := "img/car.png" 
	if !facingRight {
		imagePath = "img/car2.png"
	}

	img := canvas.NewImageFromFile(imagePath)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(80, 50)) 

	return &VehicleView{
		Vehicle: vehicle,
		Image:   img,
	}
}

func (v *VehicleView) Render() fyne.CanvasObject {
	return v.Image
}
