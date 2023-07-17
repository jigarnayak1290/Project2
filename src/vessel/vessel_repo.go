package vessel

type VesselRepo interface {
	GetAllVessel() ([]VesselData, error)
}


type DBVesselRepo struct {

}

type (dbVRepo DBVesselRepo) GetAllVessel() ([]VesselData, error) {

	vessels := []VesselData {
		VesselData{ID:1, Naccs_code: "n1", Vessel_name: "v1", Owner_name:"o1", Modified_person_name: "m1"},
		VesselData{ID:1, Naccs_code: "n2", Vessel_name: "v3", Owner_name:"o2", Modified_person_name: "m2"},
	}

	return vessels, nil
}