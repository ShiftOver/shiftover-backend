package dto

import "time"

// UserEntity represents the user entity
type UserEntity struct {
	UserID            string    `json:"userId" bson:"userId"`
	NurseID           string    `json:"nurseId" bson:"nurseId"`
	FirstName         string    `json:"firstName" bson:"firstName"`
	LastName          string    `json:"lastName" bson:"lastName"`
	Email             string    `json:"email" bson:"email"`
	ProfilePictureURL string    `json:"profilePictureUrl" bson:"profilePictureUrl,omitempty"`
	WardID            string    `json:"wardId" bson:"wardId"`
	StartDate         string    `json:"startDate" bson:"startDate"`
	DateOfBirth       string    `json:"dateOfBirth" bson:"dateOfBirth"`
	Position          string    `json:"position" bson:"position"`
	Contact           string    `json:"contact" bson:"contact"`
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt" bson:"updatedAt"`
}
