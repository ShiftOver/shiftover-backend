package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (r *userRepository) Fetch(ctx context.Context, userID string) (*dto.UserEntity, error) {
	user := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "userId", Value: userID}})

	var entity dto.UserEntity
	if err := user.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *userRepository) Exists(ctx context.Context, userID string) (bool, error) {
	user := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "userId", Value: userID}})

	var entity dto.UserEntity
	if err := user.Decode(&entity); err != nil {
		return false, nil
	}

	return true, nil
}
