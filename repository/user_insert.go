package repository

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Insert inserts a new user into the database
func (r *userRepository) Insert(ctx context.Context, payload dto.UserEntity) error {
	_, err := r.collection.InsertOne(ctx, payload)
	if err != nil {
		return errors.Wrap(err, "error - [userRepository.InsertUser]: unable to insert user")
	}

	return nil
}
