package service

import (
	"context"
	"fmt"
	"time"

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

func (s *service) InsertHospital(ctx context.Context, hospital *dto.HospitalEntity) error {
	// Check if the hospital name already exists
	if s.hospitalRepository.ExistsByName(ctx, hospital.HospitalName) {
		return errors.New("error - [service.InsertHospital]: hospital name already exists")
	}
	// Fetch the next hospital ID from the counter repository
	hospitalID, err := s.counterRepository.GetCurrentHospitalIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertHospital]: unable to fetch current hospital ID count")
	}

	// Increment the hospital ID counter
	err = s.counterRepository.IncrementHospitalIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertHospital]: unable to increment hospital ID count")
	}

	// Set the hospital ID
	hospital.HospitalID = fmt.Sprintf("HOSPITAL-%d", hospitalID)

	hospital.CreatedAt = time.Now()

	// Insert the hospital into the database
	err = s.hospitalRepository.Insert(ctx, hospital)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertHospital]: unable to insert hospital")
	}

	return nil
}
