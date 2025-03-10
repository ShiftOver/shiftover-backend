package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// Insert inserts a new room into the database
func (r *roomRepository) Insert(ctx context.Context, entity dto.RoomEntity) error {
	_, err := r.collection.InsertOne(ctx, entity)
	if err != nil {
		return errors.Wrap(err, "error - [roomRepository.Insert]: unable to insert entity")
	}

	return nil
}

// AddPatient updates the CurrentPatientID field in a room
func (r *roomRepository) AddPatient(ctx context.Context, entity dto.AddPatientRequest) error {
	filter := bson.D{{Key: "roomId", Value: entity.RoomID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "currentPatientId", Value: entity.PatientID},
			{Key: "updatedAt", Value: time.Now()},
		}},
	}
	opts := options.Update().SetUpsert(false)

	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return errors.Wrap(err, "error - [roomRepository.AddPatient]: unable to update room with patient ID")
	}

	return nil
}
