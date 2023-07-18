package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jigarnayak1290/Project2/src/handlers"
)

func main() {
	// fmt.Println("Hello from main")

	// VesselRepo := vessel.DBVesselRepo{}

	// service := service.NewVesselService(VesselRepo)
	// service.ListVessel()

	l := log.New(os.Stdout, "Vessel-api", log.LstdFlags)
	ph := handlers.NewVessel(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	C := make(chan os.Signal)
	signal.Notify(C, os.Interrupt)
	signal.Notify(C, os.Kill)

	sig := <-C
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
