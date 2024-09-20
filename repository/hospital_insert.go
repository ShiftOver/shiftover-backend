package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (r *hospitalRepository) Insert(ctx context.Context, entity dto.HospitalEntity) error {
	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return errors.Wrap(err, "error - [hospitalRepository.Insert]: unable to insert entity")
	}

	return nil
}
