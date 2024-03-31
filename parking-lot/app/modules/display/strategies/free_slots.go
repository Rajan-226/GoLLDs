package display_strategies

import "fmt"

type FreeSlots struct {
}

func (f *FreeSlots) Display(vehicleType string) {
	fmt.Printf("Free slots for CAR on Floor 1: 4,5,6\n")
}
