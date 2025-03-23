package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// GetCurrentWardIDCount fetches the current Ward id count
func (r *counterRepository) GetCurrentWardIDCount(ctx context.Context) (int, error) {
	result := r.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: "wardId"},
	})

	var entity dto.CounterEntity
	if err := result.Decode(&entity); err != nil {
		return -1, errors.Wrap(err, "error - [counterRepository.GetCurrentWardIDCount]: unable to decode result")
	}

	return entity.SequenceValue, nil
}

// IncrementWardIDCount increments the ward id count
func (r *counterRepository) IncrementWardIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "wardId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: 1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.IncrementWardIDCount]: unable to increment ward id count")
	}

	return nil
}

// DecrementWardIDCount decrements the ward id count
func (r *counterRepository) DecrementWardIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "wardId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: -1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.DecrementWardIDCount]: unable to decrement ward id count")
	}

	return nil
}
