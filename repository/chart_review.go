package repository

import "go.mongodb.org/mongo-driver/mongo"

type chartReviewRepository struct {
	collection mongo.Collection
}

// ChartReviewRepositoryConfig represents the configuration for chartReview repository
type ChartReviewRepositoryConfig struct {
	CollectionName string
}

// NewChartReviewRepository creates a new chartReview repository
func NewChartReviewRepository(collection mongo.Collection) ChartReviewRepository {
	return &chartReviewRepository{
		collection: collection,
	}
}
