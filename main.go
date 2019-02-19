package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Greeting represents the greeting
type Greeting struct {
	Message   string `json:"greeting"`
	Hostname  string `json:"hostname"`
	Name      string `json:"name"`
	TestField string `json:"testField"`
	Version   int32  `json:"version"`
}

func main() {
	log.Printf("Initializing software...")
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
		Message:   "Hello!",
		Hostname:  h,
		Name:      "Another Hello!",
		TestField: "Godzilla",
		Version:   1,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(greeting)
}
