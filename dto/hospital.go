package dto

import "time"

// HospitalEntity represents the hospital entity
type HospitalEntity struct {
	HospitalID           string    `json:"hospitalId" bson:"hospitalId" validate:"required"`
	HospitalName         string    `json:"hospitalName" bson:"hospitalName" validate:"required"`
	HospitalAbbreviation string    `json:"hospitalAbbreviation" bson:"hospitalAbbreviation" validate:"required"`
	CreatedAt            time.Time `json:"createdAt" bson:"createdAt" validate:"required"`
}
