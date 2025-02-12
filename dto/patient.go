package dto

import "time"

// Allergy represents the allergy entity
type Allergy struct {
	Name string `json:"name" bson:"name"`
}

// PatientEntity represents the patient entity
type PatientEntity struct {
	PatientID   string    `json:"patientId" bson:"patientId"`
	FirstName   string    `json:"firstName" bson:"firstName"`
	LastName    string    `json:"lastName" bson:"lastName"`
	DateOfBirth string    `json:"dateOfBirth" bson:"dateOfBirth"`
	HN          string    `json:"hn" bson:"hn"`
	Gender      string    `json:"gender" bson:"gender" validate:"oneof=male female others"`
	Allergies   []Allergy `json:"allergies" bson:"allergies"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
