package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

// Greeting represents the greeting
type Greeting struct {
	Message  string `json:"greeting"`
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
}

func main() {
	log.Printf("Initializing software...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	router := mux.NewRouter()
	router.HandleFunc("/api/greeting", GetGreeting).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
	msg := <-c
	log.Printf("Ending process... %s \n", msg)
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
