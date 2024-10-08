package repository

import "go.mongodb.org/mongo-driver/mongo"

type hospitalRepository struct {
	collection mongo.Collection
}

// HospitalRepositoryConfig represents the configuration for hospital repository
type HospitalRepositoryConfig struct {
	CollectionName string
}

// NewHospitalRepository creates a new hospital repository
func NewHospitalRepository(collection mongo.Collection) HospitalRepository {
	return &hospitalRepository{
		collection: collection,
	}
}
