package repository

import "go.mongodb.org/mongo-driver/mongo"

type patientRepository struct {
	collection mongo.Collection
}

// PatientRepositoryConfig represents the configuration for example repository
type PatientRepositoryConfig struct {
	CollectionName string
}

// NewPatientRepository creates a new example repository
func NewPatientRepository(collection mongo.Collection) PatientRepository {
	return &patientRepository{
		collection: collection,
	}
}
