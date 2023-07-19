package handlers

import (
	"log"
	"net/http"

	"github.com/jigarnayak1290/Project2/src/service"
	"github.com/jigarnayak1290/Project2/src/vessel"
)

type Vessel struct {
	l *log.Logger
}

func NewVessel(l *log.Logger) *Vessel {
	return &Vessel{l}
}

func (v *Vessel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		v.l.Println("Get", r.URL)

		queryParams := r.URL.Query()

		if len(queryParams) > 1 {
			v.l.Println("More than 1 parameer is not allowed")
			http.Error(rw, "Parameter count exceed", http.StatusBadRequest)
			return

		} else if len(queryParams) == 0 {
			v.l.Println("Get all vessel")
			v.getVessels(rw, r)
			return

		} else {
			var receivedParams []string
			for param := range queryParams {
				receivedParams = append(receivedParams, param)
			}

			if receivedParams[0] != "naccs_code" {
				v.l.Println("Wrong parameter, must be naccs_code")
				http.Error(rw, "invalid Parameter received", http.StatusBadRequest)
				return
			}
			v.GetVesselByNaccsCode(queryParams.Get("naccs_code"), rw, r)
		}

		return
	}

}

func (v *Vessel) getVessels(rw http.ResponseWriter, r *http.Request) {
	v.l.Println("Http get Vessel")

	//DBrp := vessel.DBVesselRepo{}
	//serv := service.NewVesselService(DBrp)

	PostDBrp := vessel.PostGresDBVesselRepo{}
	PostDBrp.DBInit("localhost", 5432, "postgres", "mysecretpassword", "postgres")
	serv := service.NewVesselService(PostDBrp)
	serv.ListVessel()
}

func (v *Vessel) GetVesselByNaccsCode(naccs_code string, rw http.ResponseWriter, r *http.Request) {
	v.l.Println("Http get Vessel")

	// DBrp := vessel.DBVesselRepo{}
	// serv := service.NewVesselService(DBrp)
	// serv.GetVesselByNaccsCode("1")

	PostDBrp := vessel.PostGresDBVesselRepo{}
	PostDBrp.DBInit("localhost", 5432, "postgres", "mysecretpassword", "postgres")
	serv := service.NewVesselService(PostDBrp)
	serv.GetVesselByNaccsCode(naccs_code)
}
