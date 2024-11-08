// models/vehicle.go
package models

import (
    "fmt"
    "time"
)

var currentID int = 0

type Vehicle struct {
    ID      int
    Blocked bool
}

func NewVehicle() *Vehicle {
    currentID++
    return &Vehicle{ID: currentID}
}

func (v *Vehicle) Park(duration time.Duration, onComplete func()) {
    v.Blocked = true
    go func() {
        time.Sleep(duration)
        v.Blocked = false
        if onComplete != nil {
            onComplete()
        }
    }()
}

func (v *Vehicle) Block() {
    v.Blocked = true
}

func (v *Vehicle) Unblock() {
    v.Blocked = false
}

func (v *Vehicle) Manage(parkingLot *ParkingLot, updateChan chan func(), onEntry func(int), onExit func(int)) {
    for {
        spaceIndex, entered := parkingLot.Enter(v)
        if entered {
            onEntry(spaceIndex) 
            fmt.Println("Vehículo", v.ID, "ha entrado al estacionamiento.")

            v.Park(2*time.Second, func() {
                parkingLot.Exit(spaceIndex)
                onExit(spaceIndex)        
                fmt.Println("Vehículo", v.ID, "ha salido del estacionamiento.")
            })
            return
        }
        time.Sleep(1 * time.Second)
    }
}


