package dto

// SignUpReq represents the request dto for the sign up endpoint
type SignUpReq struct {
	NurseID           string `json:"nurseId" form:"nurseId" validate:"required"`
	FirstName         string `json:"firstName" form:"firstName" validate:"required"`
	LastName          string `json:"lastName" form:"lastName" validate:"required"`
	Email             string `json:"email" form:"email" validate:"required,email"`
	Password          string `json:"password" form:"password" validate:"required"`
	ProfilePictureURL string `json:"profilePictureUrl" form:"profilePictureUrl"`
	WardID            string `json:"wardId" form:"wardId" validate:"required"`
	StartDate         string `json:"startDate" form:"startDate" validate:"required"`
	DateOfBirth       string `json:"dateOfBirth" form:"dateOfBirth" validate:"required"`
	Position          string `json:"position" form:"position" validate:"required"`
	Contact           string `json:"contact" form:"contact" validate:"required"`
}

// SignUpRes represents the response dto for the sign up endpoint
type SignUpRes struct {
	UserID  string `json:"userId"`
	NurseID string `json:"nurseId"`
	Email   string `json:"email"`
}
