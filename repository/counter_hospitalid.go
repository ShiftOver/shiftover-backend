package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// GetCurrentHospitalIDCount fetches the current Hospital id count
func (r *counterRepository) GetCurrentHospitalIDCount(ctx context.Context) (int, error) {
	result := r.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: "hospitalId"},
	})

	var entity dto.CounterEntity
	if err := result.Decode(&entity); err != nil {
		return -1, errors.Wrap(err, "error - [counterRepository.GetCurrentHospitalIDCount]: unable to decode result")
	}

	return entity.SequenceValue, nil
}

// IncrementHospitalIDCount increments the hospital id count
func (r *counterRepository) IncrementHospitalIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "hospitalId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: 1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.IncrementHospitalIDCount]: unable to increment hospital id count")
	}

	return nil
}

// DecrementHospitalIDCount decrements the hospital id count
func (r *counterRepository) DecrementHospitalIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "hospitalId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: -1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.DecrementHospitalIDCount]: unable to decrement hospital id count")
	}

	return nil
}
