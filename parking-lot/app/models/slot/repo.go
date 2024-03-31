package slot

import "sync"

var (
	slots     = make(map[int64]*Slot)
	mux       = sync.RWMutex{}
	idCounter = int64(0)
)

type IRepo interface {
	Create(slot *Slot)
}

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Create(slot *Slot) {
	mux.Lock()

	slot.SetID(idCounter)
	slots[idCounter] = slot
	idCounter++

	mux.Unlock()
}
