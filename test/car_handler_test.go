// test/car_handler_test.go
package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretch/testify/assert"
	"github.com/example.com/cars/api"
	"github.com/example.com/cars/model"
)

func TestCarHandler(t *testing.T) {
	// Create a new router
	r := mux.NewRouter()
	api.InitCarRoutes(r)

	// Test CreateCar
	t.Run("CreateCar", testCreateCar)

	// Test GetCars
	t.Run("GetCars", testGetCars)

	// Test UpdateCar
	t.Run("UpdateCar", testUpdateCar)

	// Test DeleteCar
	t.Run("DeleteCar", testDeleteCar)
}

func testCreateCar(t *testing.T) {
	// Prepare a request body with a new car
	newCar := model.Car{
		ID:    "1",
		Model: "TestCar",
	}
	requestBody, err := json.Marshal(newCar)
	assert.NoError(t, err)

	// Create a request with the prepared body
	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	// Record the HTTP response
	rec := httptest.NewRecorder()
	api.InitCarRoutes(mux.NewRouter())
	api.CreateCar(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var createdCar model.Car
	err = json.NewDecoder(rec.Body).Decode(&createdCar)
	assert.NoError(t, err)
	assert.Equal(t, newCar, createdCar)
}

func testGetCars(t *testing.T) {
	// Create a request
	req, err := http.NewRequest("GET", "/cars", nil)
	assert.NoError(t, err)

	// Record the HTTP response
	rec := httptest.NewRecorder()
	api.InitCarRoutes(mux.NewRouter())
	api.GetCars(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	var cars []model.Car
	err = json.NewDecoder(rec.Body).Decode(&cars)
	assert.NoError(t, err)
}

func testUpdateCar(t *testing.T) {
	// Create a request body with updated car information
	updatedCar := model.Car{
		ID:    "1",
		Model: "UpdatedTestCar",
	}
	requestBody, err := json.Marshal(updatedCar)
	assert.NoError(t, err)

	// Create a request with the prepared body
	req, err := http.NewRequest("PUT", "/cars/1", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	// Record the HTTP response
	rec := httptest.NewRecorder()
	api.InitCarRoutes(mux.NewRouter())
	api.UpdateCar(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Updated car with ID 1")
}

func testDeleteCar(t *testing.T) {
	// Create a request
	req, err := http.NewRequest("DELETE", "/cars/1", nil)
	assert.NoError(t, err)

	// Record the HTTP response
	rec := httptest.NewRecorder()
	api.InitCarRoutes(mux.NewRouter())
	api.DeleteCar(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	assert.Contains(t, rec.Body.String(), "Deleted car with ID 1")
}
