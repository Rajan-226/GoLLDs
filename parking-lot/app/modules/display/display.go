package display

import (
	"errors"
	"main/app/constants"
	displaystrategies "main/app/modules/display/strategies"
)

type Processor struct {
}

func NewProcessor() *Processor {
	return &Processor{}
}

// Question: what will happen if we need to pass different inputs also to display methods for implementation of different strategies?
// 1. Pass an interface as an input?
// 2. Create different method for different input? - all structs will have to implement that method
type IDisplay interface {
	Display(vehicleType string)
}

var (
	strategies = map[string]IDisplay{
		constants.FREE_COUNT:     &displaystrategies.FreeCount{},
		constants.FREE_SLOTS:     &displaystrategies.FreeSlots{},
		constants.OCCUPIED_SLOTS: &displaystrategies.OccupiedSlots{},
	}
)

func (p *Processor) Display(displayType string, vehicleType string) error {
	if _, found := strategies[displayType]; !found {
		return errors.New("no such display type")
	}

	strategies[displayType].Display(vehicleType)

	return nil
}
