package repository

import "go.mongodb.org/mongo-driver/mongo"

type wardRepository struct {
	collection mongo.Collection
}

// WardRepositoryConfig represents the configuration for example repository
type WardRepositoryConfig struct {
	CollectionName string
}

// NewWardRepository creates a new example repository
func NewWardRepository(collection mongo.Collection) WardRepository {
	return &wardRepository{
		collection: collection,
	}
}
