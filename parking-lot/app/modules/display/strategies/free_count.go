package display_strategies

import "fmt"

type FreeCount struct {
}

func (f *FreeCount) Display(vehicleType string) {

	fmt.Printf("No. of free slots for %s on Floor %d: %d\n", vehicleType, 5, 2)
}
