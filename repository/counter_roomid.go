package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// GetCurrentRoomIDCount fetches the current room id count
func (r *counterRepository) GetCurrentRoomIDCount(ctx context.Context) (int, error) {
	result := r.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: "roomId"},
	})

	var entity dto.CounterEntity
	if err := result.Decode(&entity); err != nil {
		return -1, errors.Wrap(err, "error - [counterRepository.GetCurrentRoomIDCount]: unable to decode result")
	}

	return entity.SequenceValue, nil
}

// IncrementRoomIDCount increments the room id count
func (r *counterRepository) IncrementRoomIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "roomId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: 1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.IncrementRoomIDCount]: unable to increment room id count")
	}

	return nil
}

// DecrementRoomIDCount decrements the room id count
func (r *counterRepository) DecrementRoomIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "roomId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: -1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.DecrementRoomIDCount]: unable to decrement room id count")
	}

	return nil
}
