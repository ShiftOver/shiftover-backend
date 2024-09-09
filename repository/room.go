package repository

import "go.mongodb.org/mongo-driver/mongo"

type roomRepository struct {
	collection mongo.Collection
}

// RoomRepositoryConfig represents the configuration for room repository
type RoomRepositoryConfig struct {
	CollectionName string
}

// NewRoomRepository creates a new room repository
func NewRoomRepository(collection mongo.Collection) RoomRepository {
	return &roomRepository{
		collection: collection,
	}
}
