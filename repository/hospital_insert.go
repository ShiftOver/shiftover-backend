package repository

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Insert inserts a new ward into the database
func (r *hospitalRepository) Insert(ctx context.Context, payload *dto.HospitalEntity) error {
	_, err := r.collection.InsertOne(ctx, payload)
	if err != nil {
		return errors.Wrap(err, "error - [hospitalRepository.Insert]: unable to insert hospital")
	}

	return nil
}
