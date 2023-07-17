package main

import (
	"fmt"

	"github.com/jigarnayak1290/Project2/src/service"
	"github.com/jigarnayak1290/Project2/src/vessel"
)

func main() {
	fmt.Println("Hello from main")

	VesselRepo := vessel.DBVesselRepo{}

	service := service.NewVesselService(VesselRepo)
	service.ListVessel()

}
