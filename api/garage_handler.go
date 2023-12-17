// api/garage_handler.go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/example.com/cars/model"
)



// CreateGarageEntry handles the creation of a new garage entry
func CreateGarageEntry(w http.ResponseWriter, r *http.Request) {
	var newEntry model.GarageEntry
	err := json.NewDecoder(r.Body).Decode(&newEntry)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEntry)
}

// GetGarageEntries handles fetching the list of garage entries
func GetGarageEntries(w http.ResponseWriter, r *http.Request) {

	entries := []model.GarageEntry{
		{ID: "1", CarID: "1", Status: "In"},
		{ID: "2", CarID: "2", Status: "Out"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

// UpdateGarageEntry handles updating an existing garage entry
func UpdateGarageEntry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entryID := params["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated garage entry with ID %s", entryID)
}

// DeleteGarageEntry handles deleting an existing garage entry
func DeleteGarageEntry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entryID := params["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted garage entry with ID %s", entryID)
}

// Initialize garage routes
func InitGarageRoutes(r *mux.Router) {
	r.HandleFunc("/garage", CreateGarageEntry).Methods("POST")
	r.HandleFunc("/garage", GetGarageEntries).Methods("GET")
	r.HandleFunc("/garage/{id}", UpdateGarageEntry).Methods("PUT")
	r.HandleFunc("/garage/{id}", DeleteGarageEntry).Methods("DELETE")
}
