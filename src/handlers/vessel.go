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
		v.getVessels(rw, r)
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
	return
}
