// Package repository provides the repository interfaces for the server
package repository

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// UserRepository represents the repository functions for the users collection
type UserRepository interface {
	Insert(ctx context.Context, payload dto.UserEntity) error
	Fetch(ctx context.Context, userID string) (*dto.UserEntity, error)
	Exists(ctx context.Context, userID string) bool
}

// PatientRepository represents the repository functions for the patients collection
type PatientRepository interface {
	Fetch(ctx context.Context, patientID string) (*dto.PatientEntity, error)
}

// HospitalRepository represents the repository functions for the hospitals collection
type HospitalRepository interface {
	Fetch(ctx context.Context, hospitalID string) (*dto.HospitalEntity, error)
	Exists(ctx context.Context, hospitalID string) bool
	List(ctx context.Context) ([]*dto.HospitalEntity, error)
}

// WardRepository represents the repository functions for the wards collection
type WardRepository interface {
	Fetch(ctx context.Context, wardID string) (*dto.WardEntity, error)
	Exists(ctx context.Context, wardID string) bool
	Insert(ctx context.Context, payload dto.WardEntity) error
}

// RoomRepository represents the repository functions for the rooms collection
type RoomRepository interface {
	Fetch(ctx context.Context, roomID string) (*dto.RoomEntity, error)
}

// CounterRepository represents the repository functions for the counters collection
type CounterRepository interface {
	GetCurrentUserIDCount(ctx context.Context) (int, error)
	IncrementUserIDCount(ctx context.Context) error
	DecrementUserIDCount(ctx context.Context) error
}

// FirebaseAuthRepository represents the repository functions for the firebase auth
type FirebaseAuthRepository interface {
	SignUp(ctx context.Context, payload dto.SignUpReq, userID, hospitalID string) error
}

// FirebaseStorageRepository represents the repository functions for the firebase storage
type FirebaseStorageRepository interface {
}
