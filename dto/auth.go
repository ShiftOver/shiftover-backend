package dto

// SignUpReq represents the request dto for the sign up endpoint
type SignUpReq struct {
	NurseID           string `json:"nurse_id" form:"nurse_id" validate:"required"`
	FirstName         string `json:"first_name" form:"first_name" validate:"required"`
	LastName          string `json:"last_name" form:"last_name" validate:"required"`
	Email             string `json:"email" form:"email" validate:"required,email"`
	Password          string `json:"password" form:"password" validate:"required"`
	ProfilePictureURL string `json:"profile_picture_url" form:"profile_picture_url"`
	WardID            string `json:"ward_id" form:"ward_id" validate:"required"`
	StartDate         string `json:"start_date" form:"start_date" validate:"required"`
	DateOfBirth       string `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Position          string `json:"position" form:"position" validate:"required"`
	Contact           string `json:"contact" form:"contact" validate:"required"`
}

// SignUpRes represents the response dto for the sign up endpoint
type SignUpRes struct {
	UserID      string `json:"user_id"`
	NurseID     string `json:"nurse_id"`
	Email       string `json:"email"`
	FirebaseUID string `json:"firebase_uid"`
}
