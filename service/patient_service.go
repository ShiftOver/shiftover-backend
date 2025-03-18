package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// GetPatient fetches a user by their ID
func (s *service) GetPatient(ctx context.Context, patientID string) (*dto.PatientEntity, error) {
	patient, err := s.patientRepository.Fetch(ctx, patientID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetPatient]: unable to fetch Patient")
	}
	return patient, nil
}

// InsertPatient inserts a new patient into the database
func (s *service) InsertPatient(ctx context.Context, patientRequest dto.CreatePatientRequest) error {
	patient := dto.PatientEntity{
		FirstName:            patientRequest.PatientModel.FirstName,
		LastName:             patientRequest.PatientModel.LastName,
		DateOfBirth:          patientRequest.PatientModel.DateOfBirth,
		HN:                   patientRequest.PatientModel.HN,
		Sex:                  patientRequest.PatientModel.Sex,
		Allergies:            patientRequest.PatientModel.Allergies,
		ProfilePictureURL:    patientRequest.PatientModel.ProfilePictureURL,
		Education:            patientRequest.PatientModel.Education,
		Occupation:           patientRequest.PatientModel.Occupation,
		Height:               patientRequest.PatientModel.Height,
		Weight:               patientRequest.PatientModel.Weight,
		ModeOfArrival:        patientRequest.PatientModel.ModeOfArrival,
		AdmittedForm:         patientRequest.PatientModel.AdmittedForm,
		InitialVitalSigns:    patientRequest.PatientModel.InitialVitalSigns,
		Diagnosis:            patientRequest.PatientModel.Diagnosis,
		ChiefComplaint:       patientRequest.PatientModel.ChiefComplaint,
		PastIllness:          patientRequest.PatientModel.PastIllness,
		PastIllnessHistory:   patientRequest.PatientModel.PastIllnessHistory,
		FamilyIllnessHistory: patientRequest.PatientModel.FamilyIllnessHistory,
		Reactions:            patientRequest.PatientModel.Reactions,
		Tobacco:              patientRequest.PatientModel.Tobacco,
		Alcohol:              patientRequest.PatientModel.Alcohol,
		Drugs:                patientRequest.PatientModel.Drugs,
		Exercise:             patientRequest.PatientModel.Exercise,
		Sleep:                patientRequest.PatientModel.Sleep,
		Information:          patientRequest.PatientModel.Information,
	}
	// Fetch the next patient ID from the counter repository
	patientID, err := s.counterRepository.GetCurrentPatientIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertPatient]: unable to fetch current patient ID count")
	}

	// Increment the patient ID counter
	err = s.counterRepository.IncrementPatientIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertPatient]: unable to increment patient ID count")
	}

	// Set the patient ID
	patient.PatientID = fmt.Sprintf("PATIENT-%d", patientID)

	// Set the CreatedAt field to the current time
	var timeNow = time.Now()
	patient.CreatedAt = timeNow
	patient.UpdatedAt = timeNow
	// Insert the patient into the database
	err = s.patientRepository.Insert(ctx, patient)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertPatient]: unable to insert patient")
	}

	// Add the patient to the room
	addPatienttoRoomRequest := dto.AddPatientRequest{
		RoomID:    patientRequest.RoomID,
		PatientID: patient.PatientID,
	}
	err = s.roomRepository.AddPatient(ctx, addPatienttoRoomRequest)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertPatient]: unable to add patient to room")
	}

	// Create a new chart review for the patient
	chartReview := dto.ChartReviewEntity{
		PatientID: patient.PatientID,
		Charts:    []dto.Chart{},
	}
	err = s.chartReviewRepository.Upsert(ctx, chartReview)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertPatient]: unable to create chart review")
	}

	return nil
}
