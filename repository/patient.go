package repository

import "go.mongodb.org/mongo-driver/mongo"

type patientRepository struct {
	collection mongo.Collection
}

// PatientRepositoryConfig represents the configuration for patient repository
type PatientRepositoryConfig struct {
	CollectionName string
}

// NewPatientRepository creates a new patient repository
func NewPatientRepository(collection mongo.Collection) PatientRepository {
	return &patientRepository{
		collection: collection,
	}
}
