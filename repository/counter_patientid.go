package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// GetCurrentPatientIDCount fetches the current Patient id count
func (r *counterRepository) GetCurrentPatientIDCount(ctx context.Context) (int, error) {
	result := r.collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: "patientId"},
	})

	var entity dto.CounterEntity
	if err := result.Decode(&entity); err != nil {
		return -1, errors.Wrap(err, "error - [counterRepository.GetCurrentPatientIDCount]: unable to decode result")
	}

	return entity.SequenceValue, nil
}

// IncrementPatientIDCount increments the patient id count
func (r *counterRepository) IncrementPatientIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "patientId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: 1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.IncrementPatientIDCount]: unable to increment patient id count")
	}

	return nil
}

// DecrementPatientIDCount decrements the patient id count
func (r *counterRepository) DecrementPatientIDCount(ctx context.Context) error {
	_, err := r.collection.UpdateOne(ctx, bson.D{
		{Key: "_id", Value: "patientId"},
	}, bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "sequence_value", Value: -1},
		}},
	})
	if err != nil {
		return errors.Wrap(err, "error - [counterRepository.DecrementPatientIDCount]: unable to decrement patient id count")
	}

	return nil
}
