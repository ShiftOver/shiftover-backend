package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// GetPatientAssessment fetches a hospital by their ID
func (s *service) GetPatientAssessment(ctx context.Context, patientID string) (*dto.PatientAssessmentEntity, error) {
	patientAssessment, err := s.patientAssessmentRepository.Fetch(ctx, patientID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetHospital]: unable to fetch hospital")
	}
	return patientAssessment, nil
}

// UpsertPatientAssessment upserts a patient assessment
func (s *service) UpsertPatientAssessment(ctx context.Context, patientAssessment dto.PatientAssessmentEntity) error {
	err := s.patientAssessmentRepository.Upsert(ctx, patientAssessment)
	if err != nil {
		return errors.Wrap(err, "error - [service.UpsertPatientAssessment]: unable to upsert patient assessment")
	}
	return nil
}
