// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/repository"
)

// Port represents the service layer functions
type Port interface {
	// Auth Service
	SignUp(ctx context.Context, req dto.SignUpReq) (*dto.SignUpRes, error)

	// User Service
	GetUser(ctx context.Context, userID string) (*dto.UserEntity, error)

	// Patient Service
	GetPatient(ctx context.Context, patientID string) (*dto.PatientEntity, error)
	InsertPatient(ctx context.Context, patient dto.PatientEntity) error

	// Patient Assessment Service
	GetPatientAssessment(ctx context.Context, patientID string) (*dto.PatientAssessmentEntity, error)

	// Room Service
	GetRoom(ctx context.Context, roomID string) (*dto.GetRoomResponse, error)
	ListRooms(ctx context.Context) ([]*dto.GetRoomResponse, error)

	// Hospital Service
	GetHospital(ctx context.Context, hospitalID string) (*dto.HospitalEntity, error)
	ListHospital(ctx context.Context) ([]*dto.HospitalEntity, error)
	InsertHospital(ctx context.Context, payload *dto.HospitalEntity) error

	// Medication Service
	InsertMedication(ctx context.Context, medication dto.MedicationEntity) error
	UpsertMedication(ctx context.Context, filter map[string]interface{}, update dto.MedicationEntity) error
}

type service struct {
	userRepository             repository.UserRepository
	patientRepository          repository.PatientRepository
	patientAssesmentRepository repository.PatientAssessmentRepository
	nurseMonitoringRepository  repository.NurseMonitoringRepository
	medicationRepository       repository.MedicationRepository
	nursingRepository          repository.NursingRepository
	hospitalRepository         repository.HospitalRepository
	wardRepository             repository.WardRepository
	roomRepository             repository.RoomRepository
	authRepository             repository.FirebaseAuthRepository
	storageRepository          repository.FirebaseStorageRepository
	counterRepository          repository.CounterRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	UserRepository              repository.UserRepository
	PatientRepository           repository.PatientRepository
	PatientAssessmentRepository repository.PatientAssessmentRepository
	NurseMonitoringRepository   repository.NurseMonitoringRepository
	MedicationRepository        repository.MedicationRepository
	NursingRepository           repository.NursingRepository
	HospitalRepository          repository.HospitalRepository
	WardRepository              repository.WardRepository
	RoomRepository              repository.RoomRepository
	AuthRepository              repository.FirebaseAuthRepository
	StorageRepository           repository.FirebaseStorageRepository
	CounterRepository           repository.CounterRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		userRepository:             d.UserRepository,
		patientRepository:          d.PatientRepository,
		patientAssesmentRepository: d.PatientAssessmentRepository,
		nurseMonitoringRepository:  d.NurseMonitoringRepository,
		medicationRepository:       d.MedicationRepository,
		nursingRepository:          d.NursingRepository,
		hospitalRepository:         d.HospitalRepository,
		wardRepository:             d.WardRepository,
		roomRepository:             d.RoomRepository,
		authRepository:             d.AuthRepository,
		storageRepository:          d.StorageRepository,
		counterRepository:          d.CounterRepository,
	}
}
