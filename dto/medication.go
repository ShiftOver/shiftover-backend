package dto

import "time"

// Medication represents a medication entity
type Medication struct {
	Medication string    `json:"medication" bson:"medication"`
	Dose       string    `json:"dose" bson:"dose"`
	Route      string    `json:"route,omitempty" bson:"route,omitempty"`
	Frequency  time.Time `json:"frequency" bson:"frequency"`
}

type CreateMedicationRequest struct {
	PatientID     string       `json:"patientId" bson:"patientId" validate:"required"`
	Injection     []Medication `json:"injection" bson:"injection"`
	Intravenous   []Medication `json:"intravenous" bson:"intravenous"`
	Oral          []Medication `json:"oral" bson:"oral"`
	Topical       []Medication `json:"topical" bson:"topical"`
	Drop          []Medication `json:"drop" bson:"drop"`
	Implant       []Medication `json:"implant" bson:"implant"`
	Suppositories []Medication `json:"suppositories" bson:"suppositories"`
}

// MedicationEntity represents the medication entity
type MedicationEntity struct {
	PatientID     string       `json:"patientId" bson:"patientId"`
	Injection     []Medication `json:"injection" bson:"injection"`
	Intravenous   []Medication `json:"intravenous" bson:"intravenous"`
	Oral          []Medication `json:"oral" bson:"oral"`
	Topical       []Medication `json:"topical" bson:"topical"`
	Drop          []Medication `json:"drop" bson:"drop"`
	Implant       []Medication `json:"implant" bson:"implant"`
	Suppositories []Medication `json:"suppositories" bson:"suppositories"`
	CreatedAt     time.Time    `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt" bson:"updatedAt"`
}
