package handler

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

// Get handles the GET request for the person resource.
func Get(w http.ResponseWriter, r *http.Request) {
	// p is the person to be returned
	p := Person{
		Name: "Victoria",
		Age:  29,
	}

	// set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// set the status code to OK
	w.WriteHeader(http.StatusOK)
	// encode the person into JSON and write it to the response writer
	err := json.NewEncoder(w).Encode(&p)
	if err != nil {
		// return an internal server error if there is an error encoding the person
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
