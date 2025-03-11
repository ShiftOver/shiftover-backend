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
