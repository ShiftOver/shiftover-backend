package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// Fetch fetches a patient assessment info by patient ID
func (r *patientAssessmentRepository) Fetch(ctx context.Context, patientID string) (*dto.PatientAssessmentEntity, error) {
	patientAssessment := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "patientId", Value: patientID}})

	var entity dto.PatientAssessmentEntity
	if err := patientAssessment.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [patientAssessmentRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}
