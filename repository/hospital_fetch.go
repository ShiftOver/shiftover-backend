package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (r *hospitalRepository) Fetch(ctx context.Context, hospitalID string) (*dto.HospitalEntity, error) {
	hospital := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "hospitalId", Value: hospitalID}})

	var entity dto.HospitalEntity
	if err := hospital.Decode(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *hospitalRepository) Exists(ctx context.Context, hospitalID string) (bool, error) {
	hospital := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "hospitalId", Value: hospitalID}})

	var entity dto.HospitalEntity
	if err := hospital.Decode(&entity); err != nil {
		return false, nil
	}

	return true, nil
}

func (r *hospitalRepository) List(ctx context.Context) ([]*dto.HospitalEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	results := make([]*dto.HospitalEntity, 0)
	if err := cursor.All(ctx, results); err != nil {
		return nil, err
	}

	return results, nil
}
