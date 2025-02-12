package repository

import "go.mongodb.org/mongo-driver/mongo"

type medicationRepository struct {
	collection mongo.Collection
}

// MedicationRepositoryConfig represents the configuration for medication repository
type MedicationRepositoryConfig struct {
	CollectionName string
}

// NewMedicationRepository creates a new medication repository
func NewMedicationRepository(collection mongo.Collection) MedicationRepository {
	return &medicationRepository{
		collection: collection,
	}
}
