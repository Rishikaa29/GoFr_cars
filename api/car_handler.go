// api/car_handler.go
package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/example.com/cars/model"
)

// CreateCar handles the creation of a new car entry
func CreateCar(w http.ResponseWriter, r *http.Request) {
	var newCar model.Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

// GetCars handles fetching the list of cars
func GetCars(w http.ResponseWriter, r *http.Request) {
	
	cars := []model.Car{
		{ID: "1", Model: "CarModel1"},
		{ID: "2", Model: "CarModel2"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// UpdateCar handles updating an existing car entry
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carID := params["id"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated car with ID %s", carID)
}

// DeleteCar handles deleting an existing car entry
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carID := params["id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted car with ID %s", carID)
}

// Initialize car routes
func InitCarRoutes(r *mux.Router) {
	r.HandleFunc("/cars", CreateCar).Methods("POST")
	r.HandleFunc("/cars", GetCars).Methods("GET")
	r.HandleFunc("/cars/{id}", UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")
}
