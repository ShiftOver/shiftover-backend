package repository

import "go.mongodb.org/mongo-driver/mongo"

type userRepository struct {
	collection mongo.Collection
}

// UserRepositoryConfig represents the configuration for user repository
type UserRepositoryConfig struct {
	CollectionName string
}

// NewUserRepository creates a new user repository
func NewUserRepository(collection mongo.Collection) UserRepository {
	return &userRepository{
		collection: collection,
	}
}
