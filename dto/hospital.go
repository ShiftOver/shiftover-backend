package dto

import "time"

// HospitalEntity represents the hospital entity
type HospitalEntity struct {
	HospitalID           string    `json:"hospitalId" bson:"hospitalId"`
	HospitalName         string    `json:"hospitalName" bson:"hospitalName"`
	HospitalAbbreviation string    `json:"hospitalAbbreviation" bson:"hospitalAbbreviation"`
	CreatedAt            time.Time `json:"createdAt" bson:"createdAt"`
}
