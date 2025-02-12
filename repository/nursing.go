package repository

import "go.mongodb.org/mongo-driver/mongo"

type nursingRepository struct {
	collection mongo.Collection
}

// NursingRepositoryConfig represents the configuration for nursing repository
type NursingRepositoryConfig struct {
	CollectionName string
}

// NewNursingRepository creates a new nursing repository
func NewNursingRepository(collection mongo.Collection) NursingRepository {
	return &nursingRepository{
		collection: collection,
	}
}
