package dto

import "time"

// Allergy represents the allergy entity
type Allergy struct {
	Name string `json:"name" bson:"name"`
}

type PatientModel struct {
	FirstName            string       `json:"firstName"`
	LastName             string       `json:"lastName"`
	DateOfBirth          string       `json:"dateOfBirth"`
	HN                   string       `json:"hn"`
	Sex                  string       `json:"sex" validate:"oneof=male female other"`
	Allergies            []Allergy    `json:"allergies"`
	ProfilePictureURL    string       `json:"profilePictureURL"`
	Education            string       `json:"education"`
	Occupation           string       `json:"occupation"`
	Height               float64      `json:"height"`
	Weight               float64      `json:"weight"`
	ModeOfArrival        string       `json:"modeOfArrival"`
	AdmittedForm         string       `json:"admittedForm"`
	InitialVitalSigns    VitalSigns   `json:"initialVitalSigns"`
	Diagnosis            string       `json:"diagnosis"`
	ChiefComplaint       string       `json:"chiefComplaint"`
	PastIllness          string       `json:"pastIllness"`
	PastIllnessHistory   string       `json:"pastIllnessHistory"`
	FamilyIllnessHistory string       `json:"familyIllnessHistory"`
	Reactions            []Reaction   `json:"reactions"`
	Tobacco              SubstanceUse `json:"tobacco"`
	Alcohol              SubstanceUse `json:"alcohol"`
	Drugs                SubstanceUse `json:"drugs"`
	Exercise             Exercise     `json:"exercise"`
	Sleep                Sleep        `json:"sleep"`
	Information          Information  `json:"information"`
}

// CreatePatientRequest represents the request payload for creating a patient
type CreatePatientRequest struct {
	RoomID       string       `json:"roomID" bson:"roomID"`
	PatientModel PatientModel `json:"patientModel"`
}

// PatientEntity represents the patient entity
type PatientEntity struct {
	PatientID            string       `json:"patientId" bson:"patientId"`
	FirstName            string       `json:"firstName" bson:"firstName"`
	LastName             string       `json:"lastName" bson:"lastName"`
	DateOfBirth          string       `json:"dateOfBirth" bson:"dateOfBirth"`
	HN                   string       `json:"hn" bson:"hn"`
	Sex                  string       `json:"sex" bson:"sex" validate:"oneof=male female other"`
	Allergies            []Allergy    `json:"allergies" bson:"allergies"`
	ProfilePictureURL    string       `json:"profilePictureURL" bson:"profilePictureURL"`
	Education            string       `json:"education" bson:"education"`
	Occupation           string       `json:"occupation" bson:"occupation"`
	Height               float64      `json:"height" bson:"height"`
	Weight               float64      `json:"weight" bson:"weight"`
	ModeOfArrival        string       `json:"modeOfArrival" bson:"modeOfArrival"`
	AdmittedForm         string       `json:"admittedForm" bson:"admittedForm"`
	InitialVitalSigns    VitalSigns   `json:"initialVitalSigns" bson:"initialVitalSigns"`
	Diagnosis            string       `json:"diagnosis" bson:"diagnosis"`
	ChiefComplaint       string       `json:"chiefComplaint" bson:"chiefComplaint"`
	PastIllness          string       `json:"pastIllness" bson:"pastIllness"`
	PastIllnessHistory   string       `json:"pastIllnessHistory" bson:"pastIllnessHistory"`
	FamilyIllnessHistory string       `json:"familyIllnessHistory" bson:"familyIllnessHistory"`
	Reactions            []Reaction   `json:"reactions" bson:"reactions"`
	Tobacco              SubstanceUse `json:"tobacco" bson:"tobacco"`
	Alcohol              SubstanceUse `json:"alcohol" bson:"alcohol"`
	Drugs                SubstanceUse `json:"drugs" bson:"drugs"`
	Exercise             Exercise     `json:"exercise" bson:"exercise"`
	Sleep                Sleep        `json:"sleep" bson:"sleep"`
	Information          Information  `json:"information" bson:"information"`
	CreatedAt            time.Time    `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time    `json:"updatedAt" bson:"updatedAt"`
}

// VitalSigns represents the initial vital signs of the patient
type VitalSigns struct {
	Temperature     float64       `json:"temperature" bson:"temperature"`
	HeartRate       int           `json:"heartRate" bson:"heartRate"`
	RespiratoryRate int           `json:"respiratoryRate" bson:"respiratoryRate"`
	BloodPressure   BloodPressure `json:"bloodPressure" bson:"bloodPressure"`
}

// BloodPressure represents the blood pressure of the patient
type BloodPressure struct {
	Systol  int `json:"systol" bson:"systol"`
	Diastol int `json:"diastol" bson:"diastol"`
}

// Reaction represents a reaction entity
type Reaction struct {
	Data string `json:"data" bson:"data"`
}

// SubstanceUse represents the use of substances like tobacco, alcohol, and drugs
type SubstanceUse struct {
	Status         string         `json:"status" bson:"status"`
	QuitInfo       QuitInfo       `json:"quitInfo" bson:"quitInfo"`
	ContinuousInfo ContinuousInfo `json:"continuousInfo" bson:"continuousInfo"`
}

// QuitInfo represents the quit information for substances
type QuitInfo struct {
	SmokedDuration string `json:"smokedDuration" bson:"smokedDuration"`
	QuitDuration   string `json:"quitDuration" bson:"quitDuration"`
}

// ContinuousInfo represents the continuous use information for substances
type ContinuousInfo struct {
	Duration  string `json:"duration" bson:"duration"`
	Frequency string `json:"frequency" bson:"frequency"`
}

// Exercise represents the exercise information of the patient
type Exercise struct {
	Status    string `json:"status" bson:"status"`
	Frequency string `json:"frequency" bson:"frequency"`
}

// Sleep represents the sleep information of the patient
type Sleep struct {
	Amount int    `json:"amount" bson:"amount"`
	Status string `json:"status" bson:"status"`
	Helper string `json:"helper" bson:"helper"`
}

// Information represents the information provided by the patient
type Information struct {
	ProvidedBy       string           `json:"providedBy" bson:"providedBy"`
	EmergencyContact EmergencyContact `json:"emergencyContact" bson:"emergencyContact"`
}

// EmergencyContact represents the emergency contact information of the patient
type EmergencyContact struct {
	Name         string `json:"name" bson:"name"`
	Relationship string `json:"relationship" bson:"relationship"`
	PhoneNumber  string `json:"phoneNumber" bson:"phoneNumber"`
}
