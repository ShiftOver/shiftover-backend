package repository

import "go.mongodb.org/mongo-driver/mongo"

type patientAssessmentRepository struct {
	collection mongo.Collection
}

// PatientAssessmentRepositoryConfig represents the configuration for patient assessment repository
type PatientAssessmentRepositoryConfig struct {
	CollectionName string
}

// NewPatientAssessmentRepository creates a new patient assessment repository
func NewPatientAssessmentRepository(collection mongo.Collection) PatientAssessmentRepository {
	return &patientAssessmentRepository{
		collection: collection,
	}
}
