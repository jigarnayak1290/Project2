package vessel

type VesselData struct {
	ID                   int    `json:"ID"`
	Naccs_code           string `json:"naccs_code"`
	Vessel_name          string `json:"vessel_name"`
	Owner_name           string `json:"owner_name"`
	Modified_person_name string `json:"modified_person_name"`
	Notes                string `json:"notes"`
}

func (vd VesselData) GetVesselById(id string) string {
	return id
}
