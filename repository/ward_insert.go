package repository

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Insert inserts a new ward into the database
func (r *wardRepository) Insert(ctx context.Context, payload *dto.WardEntity) error {
	_, err := r.collection.InsertOne(ctx, payload)
	if err != nil {
		return errors.Wrap(err, "error - [wardRepository.Insert]: unable to insert ward")
	}

	return nil
}
