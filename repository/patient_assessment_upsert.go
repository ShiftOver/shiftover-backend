package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (r *patientAssessmentRepository) Upsert(ctx context.Context, entity dto.PatientAssessmentEntity) error {
	timeVal := time.Now()
	entity.UpdatedAt = timeVal

	filter := bson.D{{Key: "patientId", Value: entity.PatientID}}
	existing := r.collection.FindOne(ctx, filter)

	var existingEntity dto.PatientAssessmentEntity
	if err := existing.Decode(&existingEntity); err != nil && err != mongo.ErrNoDocuments {
		return errors.Wrap(err, "error - [patientAssessmentRepository.Fetch]: unable to decode result")
	}

	if existing.Err() == nil {
		entity.CreatedAt = existingEntity.CreatedAt
		update := bson.D{{Key: "$set", Value: entity}}
		opts := options.Update().SetUpsert(true)

		_, err := r.collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			return errors.Wrap(err, "error - [patientAssessmentRepository.Upsert]: unable to update patient assessment")
		}
	} else {
		entity.CreatedAt = timeVal
		update := bson.D{{Key: "$set", Value: entity}}
		opts := options.Update().SetUpsert(true)

		_, err := r.collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			return errors.Wrap(err, "error - [patientAssessmentRepository.Upsert]: unable to insert patient assessment")
		}
	}

	return nil
}
