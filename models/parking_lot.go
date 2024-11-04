package models

import (
    "sync"
)

type ParkingLot struct {
    Capacity    int
    Occupied    int
    Entrance    *sync.Mutex
    AvailableSpots chan struct{}
}

func NewParkingLot(capacity int) *ParkingLot {
    return &ParkingLot{
        Capacity:    capacity,
        Occupied:    0,
        Entrance:    &sync.Mutex{},
        AvailableSpots: make(chan struct{}, capacity),
    }
}

func (p *ParkingLot) Enter(vehicle *Vehicle) bool {
    p.Entrance.Lock()
    defer p.Entrance.Unlock()
    
    if p.Occupied < p.Capacity {
        p.Occupied++
        p.AvailableSpots <- struct{}{}
        return true
    }
    vehicle.Blocked = true
    return false
}

func (p *ParkingLot) Exit() {
    p.Entrance.Lock()
    defer p.Entrance.Unlock()
    <-p.AvailableSpots
    p.Occupied--
}
