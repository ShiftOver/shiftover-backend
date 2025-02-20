package service

import (
	"context"
	"time"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// GetMedication fetches a medication by patientID
func (s *service) GetMedication(ctx context.Context, patientID string) (*dto.MedicationEntity, error) {
	medication, err := s.medicationRepository.Fetch(ctx, patientID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [medicationService.GetMedication]: unable to fetch medication")
	}
	return medication, nil
}

func (s *service) InsertMedication(ctx context.Context, medication dto.MedicationEntity) error {
	// Set the CreatedAt and UpdatedAt fields to the current time
	timeNow := time.Now()
	medication.CreatedAt = timeNow
	medication.UpdatedAt = timeNow

	// Insert the medication into the database
	err := s.medicationRepository.Insert(ctx, medication)
	if err != nil {
		return errors.Wrap(err, "error - [medicationService.InsertMedication]: unable to insert medication")
	}

	return nil
}

func (s *service) UpsertMedication(ctx context.Context, medication dto.MedicationEntity) error {
	err := s.medicationRepository.Upsert(ctx, medication)
	if err != nil {
		return errors.Wrap(err, "error - [service.UpsertMedication]: unable to upsert patient assessment")
	}
	return nil
}
