package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Fetch fetches a user by their ID
func (r *userRepository) Fetch(ctx context.Context, userID string) (*dto.UserEntity, error) {
	user := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "userId", Value: userID}})

	var entity dto.UserEntity
	if err := user.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [userRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}

// Exists checks if a user exists by their ID
func (r *userRepository) Exists(ctx context.Context, userID string) bool {
	user := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "userId", Value: userID}})

	var entity dto.UserEntity
	if err := user.Decode(&entity); err != nil {
		return false
	}

	return true
}
