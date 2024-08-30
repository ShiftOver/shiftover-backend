package dto

// UserEntity represents the user entity
type UserEntity struct {
	Name string `json:"name" bson:"name"`
}
