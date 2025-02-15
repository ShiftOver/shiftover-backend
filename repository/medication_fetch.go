package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Fetch fetches a medication by patientID
func (r *medicationRepository) Fetch(ctx context.Context, patientID string) (*dto.MedicationEntity, error) {
	medication := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "patientId", Value: patientID}})

	var entity dto.MedicationEntity
	if err := medication.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [medicationRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}
