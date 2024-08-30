// Package repository provides the repository interfaces for the server
package repository

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// UserRepository represents the repository functions for the user collection
type UserRepository interface {
	List(ctx context.Context) ([]*dto.UserEntity, error)
}

// FirebaseAuthRepository represents the repository functions for the firebase auth
type FirebaseAuthRepository interface {
}

// FirebaseStorageRepository represents the repository functions for the firebase storage
type FirebaseStorageRepository interface {
}
