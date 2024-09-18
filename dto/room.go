package dto

// RoomEntity represents the room entity
type RoomEntity struct {
	RoomID           string `json:"roomId" bson:"roomId"`
	RoomName         string `json:"roomName" bson:"roomName"`
	WardID           string `json:"wardId" bson:"wardId"`
	CurrentPatientID string `json:"currentPatientId,omitempty" bson:"currentPatientId,omitempty"`
	UpdatedAt        string `json:"updatedAt" bson:"updatedAt"`
	CreatedAt        string `json:"createdAt" bson:"createdAt"`
}

// GetRoomResponse represents the response dto for the get room endpoint
type GetRoomResponse struct {
	RoomID         string        `json:"roomId"`
	RoomName       string        `json:"roomName"`
	WardID         string        `json:"wardId"`
	CurrentPatient PatientEntity `json:"currentPatient,omitempty"`
	UpdatedAt      string        `json:"updatedAt"`
	CreatedAt      string        `json:"createdAt"`
}
