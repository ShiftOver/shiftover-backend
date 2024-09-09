package repository

import "go.mongodb.org/mongo-driver/mongo"

type roomRepository struct {
	collection mongo.Collection
}

// RoomRepositoryConfig represents the configuration for example repository
type RoomRepositoryConfig struct {
	CollectionName string
}

// NewRoomRepository creates a new example repository
func NewRoomRepository(collection mongo.Collection) RoomRepository {
	return &roomRepository{
		collection: collection,
	}
}
