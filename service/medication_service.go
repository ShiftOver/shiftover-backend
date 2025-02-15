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

func (s *service) UpsertMedication(ctx context.Context, filter map[string]interface{}, update dto.MedicationEntity) error {
	// Set the UpdatedAt field to the current time
	update.CreatedAt = time.Now()
	update.UpdatedAt = time.Now()

	// Upsert the medication into the database
	err := s.medicationRepository.Upsert(ctx, filter, map[string]interface{}{
		"patientId":     update.PatientID,
		"injection":     update.Injection,
		"intravenous":   update.Intravenous,
		"oral":          update.Oral,
		"topical":       update.Topical,
		"drop":          update.Drop,
		"implant":       update.Implant,
		"suppositories": update.Suppositories,
		"createdAt":     update.CreatedAt,
		"updatedAt":     update.UpdatedAt,
	})
	if err != nil {
		return errors.Wrap(err, "error - [medicationService.UpsertMedication]: unable to upsert medication")
	}

	return nil
}
