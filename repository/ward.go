package repository

import "go.mongodb.org/mongo-driver/mongo"

type wardRepository struct {
	collection mongo.Collection
}

// WardRepositoryConfig represents the configuration for ward repository
type WardRepositoryConfig struct {
	CollectionName string
}

// NewWardRepository creates a new ward repository
func NewWardRepository(collection mongo.Collection) WardRepository {
	return &wardRepository{
		collection: collection,
	}
}
