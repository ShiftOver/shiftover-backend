package repository

import "go.mongodb.org/mongo-driver/mongo"

type nurseMonitoringRepository struct {
	collection mongo.Collection
}

// NurseMonitoringRepositoryConfig represents the configuration for nurse monitoring repository
type NurseMonitoringRepositoryConfig struct {
	CollectionName string
}

// NewNurseMonitoringRepository creates a new nurse monitoring repository
func NewNurseMonitoringRepository(collection mongo.Collection) NurseMonitoringRepository {
	return &nurseMonitoringRepository{
		collection: collection,
	}
}
