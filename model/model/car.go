// model/car.go
package model

// Car represents the model for a car
type Car struct {
	ID    string `json:"id"`
	Model string `json:"model"`
    Status string `json:"status"`
}
