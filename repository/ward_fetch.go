package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (r *wardRepository) Fetch(ctx context.Context, wardID string) (*dto.WardEntity, error) {
	ward := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "wardId", Value: wardID}})

	var entity dto.WardEntity
	if err := ward.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *wardRepository) Exists(ctx context.Context, wardID string) (bool, error) {
	ward := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "wardId", Value: wardID}})

	var entity dto.WardEntity
	if err := ward.Decode(&entity); err != nil {
		return false, nil
	}

	return true, nil
}
