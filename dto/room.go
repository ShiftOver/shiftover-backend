package dto

import "time"

// RoomEntity represents the room entity
type RoomEntity struct {
	RoomID           string    `json:"roomId" bson:"roomId"`
	RoomName         string    `json:"roomName" bson:"roomName"`
	WardID           string    `json:"wardId" bson:"wardId"`
	CurrentPatientID string    `json:"currentPatientId,omitempty" bson:"currentPatientId,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt" bson:"updatedAt"`
	CreatedAt        time.Time `json:"createdAt" bson:"createdAt"`
}

// GetRoomResponse represents the response dto for the get room endpoint
type GetRoomResponse struct {
	RoomID         string        `json:"roomId"`
	RoomName       string        `json:"roomName"`
	WardID         string        `json:"wardId"`
	CurrentPatient PatientEntity `json:"currentPatient,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	CreatedAt      time.Time     `json:"createdAt"`
}
