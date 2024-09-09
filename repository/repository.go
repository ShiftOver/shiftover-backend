// Package repository provides the repository interfaces for the server
package repository

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// UserRepository represents the repository functions for the users collection
type UserRepository interface {
	List(ctx context.Context) ([]*dto.UserEntity, error)
}

// PatientRepository represents the repository functions for the patients collection
type PatientRepository interface {
}

// HospitalRepository represents the repository functions for the hospitals collection
type HospitalRepository interface {
}

// WardRepository represents the repository functions for the wards collection
type WardRepository interface {
}

// RoomRepository represents the repository functions for the rooms collection
type RoomRepository interface {
}

// CounterRepository represents the repository functions for the counters collection
type CounterRepository interface {
}

// FirebaseAuthRepository represents the repository functions for the firebase auth
type FirebaseAuthRepository interface {
	SignUp(ctx context.Context, payload *dto.SignUpReq, userID, hospitalID string) error
}

// FirebaseStorageRepository represents the repository functions for the firebase storage
type FirebaseStorageRepository interface {
}
