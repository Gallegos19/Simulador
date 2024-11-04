package scenes

import (
    "fmt"
    "strconv"
    "strings" 
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "main/models"
    "main/views"
)

type SimulationScene struct {
    ParkingLot       *models.ParkingLot
    ParkingLotView   *views.ParkingLotView
    EventList        *widget.Label
    Content          *fyne.Container
    WaitingVehicles  []*models.Vehicle 
}


func NewSimulationScene(parkingLot *models.ParkingLot) *SimulationScene {
    parkingLotView := views.NewParkingLotView(parkingLot)

    scene := &SimulationScene{
        ParkingLot:      parkingLot,
        ParkingLotView:  parkingLotView,
        EventList:       widget.NewLabel("Eventos de entrada y salida:\n"),
    }

    scene.Content = container.NewVBox(
        parkingLotView.Render(),
        scene.EventList,
    )

    go scene.runSimulation()

    return scene
}

func (scene *SimulationScene) runSimulation() {
    vehicleID := 0
    for {
        time.Sleep(1 * time.Second) 


        vehicle := &models.Vehicle{ID: vehicleID}
        vehicleID++

        // Intenta ingresar el vehículo (ya sea de la cola o uno nuevo)
        if len(scene.WaitingVehicles) == 0 && scene.ParkingLot.Enter(vehicle) {
            var spaceIndex int
            for i, v := range scene.ParkingLotView.Spaces {
                if v == nil { // Encuentra un espacio vacío
                    spaceIndex = i
                    break
                }
            }

            // Crea y asigna la vista del vehículo en el espacio encontrado
            vehicleView := views.NewVehicleView(vehicle, true) 
            scene.ParkingLotView.Spaces[spaceIndex] = vehicleView
            scene.ParkingLotView.UpdateParkingLot()

            // Registro de entrada en el estacionamiento
            scene.addEvent("Vehículo " + strconv.Itoa(vehicle.ID) + " ha entrado al estacionamiento.", false)
            fmt.Println("Vehículo", vehicle.ID, "ha entrado al estacionamiento.")

            // Simula el tiempo de estacionamiento
            go func(v *models.Vehicle, vView *views.VehicleView, index int) {
                time.Sleep(30 * time.Second) // Tiempo estacionado
                scene.ParkingLot.Exit()
                scene.ParkingLotView.Spaces[index] = nil // Libera el espacio
                scene.ParkingLotView.UpdateParkingLot()

                // Si hay vehículos en la cola, procesar el siguiente
                if len(scene.WaitingVehicles) > 0 {
                    nextVehicle := scene.WaitingVehicles[0]
                    scene.WaitingVehicles = scene.WaitingVehicles[1:] // Remueve el vehículo de la cola de espera
                    scene.addEvent("Vehículo " + strconv.Itoa(nextVehicle.ID) + " tiene prioridad para entrar.", false)
                    scene.WaitingVehicles = append(scene.WaitingVehicles, nextVehicle)
                }

                // Registro de salida del estacionamiento
                scene.addEvent("Vehículo " + strconv.Itoa(v.ID) + " ha salido del estacionamiento.", false)
                fmt.Println("Vehículo", v.ID, "ha salido del estacionamiento.")
            }(vehicle, vehicleView, spaceIndex)
        } else {
            // Si no hay espacio, añade el vehículo a la cola de espera
            scene.WaitingVehicles = append(scene.WaitingVehicles, vehicle)
            // Registro de bloqueo de vehículo si no hay espacio
            scene.addEvent("Vehículo " + strconv.Itoa(vehicle.ID) + " esperando espacio.", true)
            fmt.Println("Vehículo", vehicle.ID, "esperando espacio.")
        }
    }
}



func (scene *SimulationScene) addEvent(event string, isWaiting bool) {
    // Limitar el número máximo de líneas en EventList
    maxLines := 50
    maxChars := 1000 // Limite de caracteres total en el evento

    // Actualiza el texto de la lista de eventos en la interfaz
    currentText := scene.EventList.Text
    lines := strings.Split(currentText, "\n")
    lines = append(lines, event) // Agregar el nuevo evento

    // Si excede el máximo de líneas, elimina las más antiguas
    if len(lines) > maxLines {
        lines = lines[len(lines)-maxLines:] // Mantiene solo las últimas `maxLines` líneas
    }

    // Unir las líneas y limitar la longitud total
    newText := strings.Join(lines, "\n")
    if len(newText) > maxChars {
        newText = newText[len(newText)-maxChars:] // Mantener solo los últimos `maxChars` caracteres
    }

    scene.EventList.SetText(newText)
    scene.EventList.Refresh()

    // Eliminar mensajes de entrada o salida después de 5 segundos
    if !isWaiting {
        go func() {
            time.Sleep(5 * time.Second) // Duración del mensaje temporal
            // Removemos el mensaje de entrada o salida
            currentText := scene.EventList.Text
            updatedText := removeEvent(currentText, event)
            scene.EventList.SetText(updatedText)
            scene.EventList.Refresh()
        }()
    } else {
        // Eliminar mensajes de espera después de 5 segundos
        go func() {
            time.Sleep(5 * time.Second) // Duración del mensaje de espera
            currentText := scene.EventList.Text
            updatedText := removeEvent(currentText, event)
            scene.EventList.SetText(updatedText)
            scene.EventList.Refresh()
        }()
    }
}



func removeEvent(text, event string) string {
    // Elimina el evento específico del texto actual
    lines := strings.Split(text, "\n")
    var updatedLines []string
    for _, line := range lines {
        if line != event {
            updatedLines = append(updatedLines, line)
        }
    }
    return strings.Join(updatedLines, "\n")
}

func (scene *SimulationScene) Render() fyne.CanvasObject {
    return scene.Content
}