package main

import (
	"main/app/boot"
)

func main() {
	server := boot.Init()

	server.CreateParkingLot("PR1234", 2, 6)

	server.Display("free_count", "car")
}
