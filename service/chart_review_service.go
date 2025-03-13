package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// GetChartReview fetches a chart review by patientID
func (s *service) GetChartReview(ctx context.Context, patientID string) (*dto.ChartReviewEntity, error) {
	chartReview, err := s.chartReviewRepository.Fetch(ctx, patientID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetChartReview]: unable to fetch chart review")
	}
	return chartReview, nil
}

// UpsertChartReview upserts a chart review into the database
func (s *service) UpsertChartReview(ctx context.Context, req dto.UpdateChartReviewRequest) error {
	chartReview := dto.ChartReviewEntity{
		PatientID: req.PatientID,
		Charts:    req.Charts,
	}

	err := s.chartReviewRepository.Upsert(ctx, chartReview)
	if err != nil {
		return errors.Wrap(err, "error - [service.UpsertChartReview]: unable to upsert chart review")
	}
	return nil
}
