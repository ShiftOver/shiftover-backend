package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// Insert inserts a new room into the database
func (r *roomRepository) Insert(ctx context.Context, entity dto.RoomEntity) error {
	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return errors.Wrap(err, "error - [roomRepository.Insert]: unable to insert entity")
	}

	return nil
}
