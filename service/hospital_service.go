package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// GetHospital fetches a hospital by their ID
func (s *service) GetHospital(ctx context.Context, hospitalID string) (*dto.HospitalEntity, error) {
	hospital, err := s.hospitalRepository.Fetch(ctx, hospitalID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetHospital]: unable to fetch hospital")
	}
	return hospital, nil
}

// ListHospital fetches all hospitals in the database
func (s *service) ListHospital(ctx context.Context) ([]*dto.HospitalEntity, error) {
	hospitals, err := s.hospitalRepository.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.ListHospital]: unable to list hospitals")
	}
	return hospitals, nil
}
