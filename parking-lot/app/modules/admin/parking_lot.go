package admin

import (
	"main/app/constants"
	"main/app/models/slot"
)

type Processor struct {
	slotRepo slot.IRepo
}

var (
	initialTrucks = 1
	initialBikes  = 2
)

func NewProcessor(slotRepo slot.IRepo) *Processor {
	return &Processor{
		slotRepo: slotRepo,
	}
}

func (p *Processor) CreateParkingLot(id string, floors int, slots int) {
	for floor := 0; floor < floors; floor++ {
		for slotNo := 0; slotNo < slots; slotNo++ {

			newSlot := slot.NewSlot(int64(floor), id)

			if slotNo < initialTrucks {
				newSlot.AddNewVehicle(constants.TRUCK)
			} else if slotNo < (initialBikes + initialTrucks) {
				newSlot.AddNewVehicle(constants.BIKE)
			} else {
				newSlot.AddNewVehicle(constants.CAR)
			}

			p.slotRepo.Create(newSlot)
		}
	}

}
