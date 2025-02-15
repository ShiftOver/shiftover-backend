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

func (s *service) InsertPatient(ctx context.Context, patient dto.PatientEntity) error {
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

	return nil
}
