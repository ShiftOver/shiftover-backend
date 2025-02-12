package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// Insert inserts a patient into the database
func (r *patientRepository) Insert(ctx context.Context, entity dto.PatientEntity) error {
	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return errors.Wrap(err, "error - [patientRepository.Insert]: unable to insert patient")
	}

	return nil
}
