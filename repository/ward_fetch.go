package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Fetch fetches a ward by its ID
func (r *wardRepository) Fetch(ctx context.Context, wardID string) (*dto.WardEntity, error) {
	ward := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "wardId", Value: wardID}})

	var entity dto.WardEntity
	if err := ward.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [wardRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}

// Exists checks if a ward exists by its ID
func (r *wardRepository) Exists(ctx context.Context, wardID string) bool {
	ward := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "wardId", Value: wardID}})

	var entity dto.WardEntity
	if err := ward.Decode(&entity); err != nil {
		return false
	}

	return true
}
