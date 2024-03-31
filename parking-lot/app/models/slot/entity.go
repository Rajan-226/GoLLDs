package slot

// Question - Is it ok to keep array of vechilesAllowed in slot? If not, how to handle this?
type Slot struct {
	id              int64
	reserved        bool
	floorNumber     int64
	parkingLotID    string
	vehiclesAllowed []string
}

func NewSlot(floorNumber int64, parkingLotID string) *Slot {
	return &Slot{
		reserved:        false,
		floorNumber:     floorNumber,
		parkingLotID:    parkingLotID,
		vehiclesAllowed: []string{},
	}
}

func (s *Slot) WithVehicleAllowed(vehiclesAllowed []string) *Slot {
	s.vehiclesAllowed = vehiclesAllowed
	return s
}

func (s *Slot) GetID() int64 {
	return s.id
}

func (s *Slot) SetID(ID int64) {
	s.id = ID
}

func (s *Slot) GetReserved() bool {
	return s.reserved
}

func (s *Slot) GetFloorNumber() int64 {
	return s.floorNumber
}

func (s *Slot) AddNewVehicle(newVehicle string) {
	s.vehiclesAllowed = append(s.vehiclesAllowed, newVehicle)
}

func (s *Slot) IsVehicleAllowed(curVehicle string) bool {
	for _, vehiclesAllowed := range s.vehiclesAllowed {
		if vehiclesAllowed == curVehicle {
			return true
		}
	}

	return false
}

func (s *Slot) MarkReserved() {
	s.reserved = true
}

func (s *Slot) MarkFree() {
	s.reserved = false
}
