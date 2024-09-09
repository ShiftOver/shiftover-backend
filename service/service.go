// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/repository"
)

// Port represents the service layer functions
type Port interface {
	ListUser(ctx context.Context) ([]*dto.UserEntity, error)
}

type service struct {
	userRepository     repository.UserRepository
	patientRepository  repository.PatientRepository
	hospitalRepository repository.HospitalRepository
	wardRepository     repository.WardRepository
	roomRepository     repository.RoomRepository
	authRepository     repository.FirebaseAuthRepository
	storageRepository  repository.FirebaseStorageRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	UserRepository     repository.UserRepository
	PatientRepository  repository.PatientRepository
	HospitalRepository repository.HospitalRepository
	WardRepository     repository.WardRepository
	RoomRepository     repository.RoomRepository
	AuthRepository     repository.FirebaseAuthRepository
	StorageRepository  repository.FirebaseStorageRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		userRepository:     d.UserRepository,
		patientRepository:  d.PatientRepository,
		hospitalRepository: d.HospitalRepository,
		wardRepository:     d.WardRepository,
		roomRepository:     d.RoomRepository,
		authRepository:     d.AuthRepository,
		storageRepository:  d.StorageRepository,
	}
}
