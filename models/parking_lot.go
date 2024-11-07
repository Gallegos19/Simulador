package models

import (
	"fmt"
	"sync"
)

type ParkingLot struct {
    Capacity       int
    OccupiedSpaces map[int]*Vehicle
    Entrance       *sync.Mutex
    AvailableSpots chan int
}

func NewParkingLot(capacity int) *ParkingLot {
    availableSpots := make(chan int, capacity)
    for i := 0; i < capacity; i++ {
        availableSpots <- i
    }

    return &ParkingLot{
        Capacity:       capacity,
        OccupiedSpaces: make(map[int]*Vehicle),
        Entrance:       &sync.Mutex{},
        AvailableSpots: availableSpots,
    }
}

func (p *ParkingLot) Enter(vehicle *Vehicle) (int, bool) {
    p.Entrance.Lock()
    defer p.Entrance.Unlock()

    select {
    case spaceIndex := <-p.AvailableSpots:
        p.OccupiedSpaces[spaceIndex] = vehicle
        return spaceIndex, true
    default:
        vehicle.Block()
        return -1, false
    }
}

func (p *ParkingLot) Exit(spaceIndex int) {
    p.Entrance.Lock()
    defer p.Entrance.Unlock()
    delete(p.OccupiedSpaces, spaceIndex)
    fmt.Println("Espacio liberado:", spaceIndex) 
    p.AvailableSpots <- spaceIndex
}

