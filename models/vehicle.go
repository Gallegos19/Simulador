package models

import "time"

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
