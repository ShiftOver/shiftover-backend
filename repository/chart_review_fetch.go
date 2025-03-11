package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// Fetch fetches a chart review by patientID
func (r *chartReviewRepository) Fetch(ctx context.Context, patientID string) (*dto.ChartReviewEntity, error) {
	chartReview := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "patientId", Value: patientID}})

	var entity dto.ChartReviewEntity
	if err := chartReview.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [chartReviewRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}
