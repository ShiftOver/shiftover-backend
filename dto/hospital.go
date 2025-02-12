package dto

import "time"

// HospitalModel represents the hospital model request
type HospitalModel struct {
	HospitalName         string `json:"hospitalName"  validate:"required"`
	HospitalAbbreviation string `json:"hospitalAbbreviation" validate:"required"`
}

// HospitalEntity represents the hospital entity
type HospitalEntity struct {
	HospitalID           string    `json:"hospitalId" bson:"hospitalId"`
	HospitalName         string    `json:"hospitalName" bson:"hospitalName"`
	HospitalAbbreviation string    `json:"hospitalAbbreviation" bson:"hospitalAbbreviation"`
	CreatedAt            time.Time `json:"createdAt" bson:"createdAt"`
}
