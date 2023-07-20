package vessel

import (
	"encoding/json"
	"io"
)

type VesselData struct {
	ID                   int    `json:"ID"`
	Naccs_code           string `json:"naccs_code"`
	Vessel_name          string `json:"vessel_name"`
	Owner_name           string `json:"owner_name"`
	Modified_person_name string `json:"modified_person_name"`
	Notes                string `json:"notes"`
}

func (v *VesselData) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(v)
}

func (v *VesselData) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(v)
}

func (vd VesselData) GetVesselById(id string) string {
	return id
}
