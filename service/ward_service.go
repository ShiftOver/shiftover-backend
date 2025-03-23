package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// GetWard fetches a ward by its ID
func (s *service) GetWard(ctx context.Context, wardID string) (*dto.WardEntity, error) {
	ward, err := s.wardRepository.Fetch(ctx, wardID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetWard]: unable to fetch ward")
	}
	return ward, nil
}

// ListWards fetches all wards in the database
func (s *service) ListWards(ctx context.Context) ([]*dto.WardEntity, error) {
	wards, err := s.wardRepository.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.ListWards]: unable to list wards")
	}
	return wards, nil
}

// InsertWard inserts a new ward into the database
func (s *service) InsertWard(ctx context.Context, ward dto.WardEntity) error {
	wardID, err := s.counterRepository.GetCurrentWardIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertWard]: unable to fetch current ward ID count")
	}

	// Increment the ward ID counter
	err = s.counterRepository.IncrementWardIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertWard]: unable to increment ward ID count")
	}

	// Set the ward ID
	ward.WardID = fmt.Sprintf("WARD-%d", wardID)

	ward.CreatedAt = time.Now()

	// Insert the ward into the database
	err = s.wardRepository.Insert(ctx, ward)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertWard]: unable to insert ward")
	}

	return nil
}
