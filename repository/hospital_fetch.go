package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Fetch fetches a hospital by their ID
func (r *hospitalRepository) Fetch(ctx context.Context, hospitalID string) (*dto.HospitalEntity, error) {
	hospital := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "hospitalId", Value: hospitalID}})

	var entity dto.HospitalEntity
	if err := hospital.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [hospitalRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}

// Exists checks if a hospital exists by their ID
func (r *hospitalRepository) Exists(ctx context.Context, hospitalID string) bool {
	hospital := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "hospitalId", Value: hospitalID}})

	var entity dto.HospitalEntity
	if err := hospital.Decode(&entity); err != nil {
		return false
	}

	return true
}

// List fetches all hospitals in the database
func (r *hospitalRepository) List(ctx context.Context) ([]*dto.HospitalEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "error - [hospitalRepository.List]: unable to list hospitals")
	}

	results := make([]*dto.HospitalEntity, 0)
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "error - [hospitalRepository.List]: unable to decode results")
	}

	return results, nil
}
