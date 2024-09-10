package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (r *counterRepository) GetCurrentUserIDCount(ctx context.Context) (int, error) {
	result := r.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: "userId"},
	})

	var entity dto.CounterEntity
	if err := result.Decode(&entity); err != nil {
		return -1, errors.Wrap(err, "error - [counterRepository.GetCurrentUserIDCount]: unable to decode result")
	}

	return entity.SequenceValue, nil
}

func (r *counterRepository) IncrementUserIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "userId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: 1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.IncrementUserIDCount]: unable to increment user id count")
	}

	return nil
}

func (r *counterRepository) DecrementUserIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "userId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: -1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.DecrementUserIDCount]: unable to decrement user id count")
	}

	return nil
}
