package api

import (
	"fmt"
	"main/app/modules/admin"
	"main/app/modules/display"
)

type Server struct {
	adminProcessor   admin.Processor
	displayProcessor display.Processor
}

func NewServer(adminProcessor *admin.Processor) *Server {
	return &Server{
		adminProcessor: *adminProcessor,
	}
}

func (s *Server) CreateParkingLot(id string, floors int, slots int) {
	s.adminProcessor.CreateParkingLot(id, floors, slots)

	fmt.Printf("Created parking lot with %d floors and %d slots per floor\n", floors, slots)
}

func ParkVehicle(vehicleType string, regNo string, color string) {

}

func UnParkVehicle(ticketID string) {

}

func (s *Server) Display(displayType string, vehicleType string) {
	if err := s.displayProcessor.Display(displayType, vehicleType); err != nil {
		fmt.Printf("Not able to display due to: %s\n", err.Error())
		return
	}
}
