package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// Fetch fetches a patient by patient ID
func (r *patientRepository) Fetch(ctx context.Context, patientID string) (*dto.PatientEntity, error) {
	patient := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "patientId", Value: patientID}})

	var entity dto.PatientEntity
	if err := patient.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [patientRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}
