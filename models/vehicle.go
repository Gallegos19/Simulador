package models

import "time"

type Vehicle struct {
    ID      int
    Blocked bool
}

func (v *Vehicle) Park(duration time.Duration) {
    time.Sleep(duration)
    v.Blocked = false
}
