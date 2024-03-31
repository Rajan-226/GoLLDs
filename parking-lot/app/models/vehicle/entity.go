package vehicle

type Vehicle struct {
	id             string
	registrationNo string
	color          string
	vehicleType    int64
	slotID         int64
}

func NewVehicle(registrationNo string, color string, vehicleType, slotID int64) *Vehicle {
	return &Vehicle{
		registrationNo: registrationNo,
		color:          color,
		vehicleType:    vehicleType,
		slotID:         slotID,
	}
}

func (v *Vehicle) GetID() string {
	return v.id
}

func (v *Vehicle) GetRegistrationNo() string {
	return v.registrationNo
}

func (v *Vehicle) GetColor() string {
	return v.color
}

func (v *Vehicle) GetVehicleType() int64 {
	return v.vehicleType
}

func (v *Vehicle) GetSlotID() int64 {
	return v.slotID
}

func (v *Vehicle) SetID(id int64) {
	v.slotID = id
}
