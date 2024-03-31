package vehicle

import (
	"sync"
)

var (
	vehicles  = make(map[int64]*Vehicle)
	mux       = sync.RWMutex{}
	idCounter = int64(0)
)

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func Create(vehicle *Vehicle) {
	mux.Lock()

	vehicle.SetID(idCounter)
	vehicles[idCounter] = vehicle
	idCounter++

	mux.Unlock()
}
