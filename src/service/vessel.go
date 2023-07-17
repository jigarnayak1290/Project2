package service

import (
	"fmt"

	"github.com/jigarnayak1290/Project2/src/vessel"
)

type VesselService struct {
	vesselRe vessel.VesselRepo
}

func NewVesselService(repo vessel.VesselRepo) *VesselService {
	return &VesselService{vesselRe: repo}
}

func (vs *VesselService) ListVessel() {
	vessel, err := vs.vesselRe.GetAllVessel()

	if err != nil {
		fmt.Println("Failed to retrive vessle data")
		return
	}
	for _, vessel := range vessel {
		fmt.Printf("Vessel : %d, ", vessel.ID)
	}
}
