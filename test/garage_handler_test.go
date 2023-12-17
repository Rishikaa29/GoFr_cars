// test/garage_handler_test.go
package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretch/testify/assert"
	"github.com/example.com/cars/api"
	"github.com/example.com/cars/model"
)

func TestGarageHandler(t *testing.T) {
	// Create a new router
	r := mux.NewRouter()
	api.InitGarageRoutes(r)

	// Test CreateGarageEntry
	t.Run("CreateGarageEntry", testCreateGarageEntry)

	// Test GetGarageEntries
	t.Run("GetGarageEntries", testGetGarageEntries)

	// Test UpdateGarageEntry
	t.Run("UpdateGarageEntry", testUpdateGarageEntry)

	// Test DeleteGarageEntry
	t.Run("DeleteGarageEntry", testDeleteGarageEntry)
}

func testCreateGarageEntry(t *testing.T) {
	// Prepare a request body with a new garage entry
	newEntry := model.GarageEntry{
		ID:     "1",
		CarID:  "1",
		Status: "In",
	}
	requestBody, err := json.Marshal(newEntry)
	assert.NoError(t, err)

	// Create a request with the prepared body
	req, err := http.NewRequest("POST", "/garage", bytes.NewBuffer(requestBody))
	assert.NoError(t, err)

	// Record the HTTP response
	rec := httptest.NewRecorder()

	// Serve the HTTP request to the recorder
	api.InitGarageRoutes(mux.NewRouter())
	api.CreateGarageEntry(rec, req)

	// Assert the HTTP response status code
	assert.Equal
