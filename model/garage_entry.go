// model/garage_entry.go
package model

// GarageEntry represents the model for a garage entry
type GarageEntry struct {
	ID     string `json:"id"`
	CarID  string `json:"car_id"`
	Status string `json:"status"` // In or Out
	
}
