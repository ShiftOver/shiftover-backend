package repository

import "go.mongodb.org/mongo-driver/mongo"

type counterRepository struct {
	collection mongo.Collection
}

// CounterRepositoryConfig represents the configuration for counter repository
type CounterRepositoryConfig struct {
	CollectionName string
}

// NewCounterRepository creates a new counter repository
func NewCounterRepository(collection mongo.Collection) CounterRepository {
	return &counterRepository{
		collection: collection,
	}
}
