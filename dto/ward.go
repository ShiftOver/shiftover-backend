package dto

import "time"

// WardEntity represents the ward entity
type WardEntity struct {
	WardID     string    `json:"wardId" bson:"wardId"`
	WardName   string    `json:"wardName" bson:"wardName"`
	HospitalID string    `json:"hospitalId" bson:"hospitalId"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}
