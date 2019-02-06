package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	rpio "github.com/stianeikeland/go-rpio"
)

// Greeting represents the greeting
type Greeting struct {
	Message  string `json:"greeting"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
}

var pin rpio.Pin
var sleepTime time.Duration

func main() {
	log.Printf("Initializing software...")
	err := rpio.Open()
	sleepTime, _ = time.ParseDuration("5s")
	if err != nil {
		log.Printf("An error occurred while trying to connect GPIO pins: %s", err.Error())
	} else {
		pin = rpio.Pin(17)
		defer rpio.Close()
		pin.Output()
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/greeting", GetGreeting).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetGreeting responses with a greeting in json form.
func GetGreeting(w http.ResponseWriter, r *http.Request) {
	log.Println("Received greeting request.")
	pin.High()
	time.Sleep(sleepTime)
	pin.Low()
	log.Println("Shut down LED.")
	h, err := os.Hostname()
	if err != nil {
		log.Fatal("Could not get hostname!")
		w.WriteHeader(500)
	}
	greeting := Greeting{
		Message:  "Hello!",
		Hostname: h,
		Name:     "Another Hello!",
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(greeting)
}
