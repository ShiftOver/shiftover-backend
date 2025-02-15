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
	Insert(ctx context.Context, entity dto.PatientEntity) error
}

// PatientAssessmentRepository represents the repository functions for the patient assessments collection
type PatientAssessmentRepository interface {
	Fetch(ctx context.Context, patientID string) (*dto.PatientAssessmentEntity, error)
}

// NurseMonitoringRepository represents the repository functions for the nurse monitoring collection
type NurseMonitoringRepository interface {
}

// MedicationRepository represents the repository functions for the medications collection
type MedicationRepository interface {
}

// NursingRepository represents the repository functions for the nursing collection
type NursingRepository interface {
}

// HospitalRepository represents the repository functions for the hospitals collection
type HospitalRepository interface {
	Fetch(ctx context.Context, hospitalID string) (*dto.HospitalEntity, error)
	Exists(ctx context.Context, hospitalID string) bool
	ExistsByName(ctx context.Context, hospitalName string) bool
	List(ctx context.Context) ([]*dto.HospitalEntity, error)
	Insert(ctx context.Context, payload *dto.HospitalEntity) error
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
	List(ctx context.Context) ([]*dto.RoomEntity, error)
	Insert(ctx context.Context, entity dto.RoomEntity) error
}

// CounterRepository represents the repository functions for the counters collection
type CounterRepository interface {
	GetCurrentUserIDCount(ctx context.Context) (int, error)
	IncrementUserIDCount(ctx context.Context) error
	DecrementUserIDCount(ctx context.Context) error
	GetCurrentHospitalIDCount(ctx context.Context) (int, error)
	IncrementHospitalIDCount(ctx context.Context) error
	DecrementHospitalIDCount(ctx context.Context) error
}

// FirebaseAuthRepository represents the repository functions for the firebase auth
type FirebaseAuthRepository interface {
	SignUp(ctx context.Context, payload dto.SignUpReq, userID, hospitalID string) error
}

// FirebaseStorageRepository represents the repository functions for the firebase storage
type FirebaseStorageRepository interface {
}
