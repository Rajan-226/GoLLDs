package display_strategies

import "fmt"

type OccupiedSlots struct {
}

func (f *OccupiedSlots) Display(vehicleType string) {
	fmt.Printf("Occupied slots for CAR on Floor 1:")
}
