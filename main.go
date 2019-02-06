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

func main() {
	err := rpio.Open()
	if err != nil {
		log.Printf("An error occurred while trying to connect GPIO pins: %s", err.Error())
	} else {
		pin := rpio.Pin(17)
		defer rpio.Close()
		pin.Output()
		for i := 0; i < 25; i++ {
			pin.Toggle()
			time.Sleep(time.Second)
		}
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/greeting", GetGreeting).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetGreeting responses with a greeting in json form.
func GetGreeting(w http.ResponseWriter, r *http.Request) {
	log.Println("Received greeting request.")
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
