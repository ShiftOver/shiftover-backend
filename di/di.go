// Package di provides dependency injection for the server
package di

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/ShiftOver/shiftover-backend/config"
	"github.com/ShiftOver/shiftover-backend/handler"
	"github.com/ShiftOver/shiftover-backend/repository"
	"github.com/ShiftOver/shiftover-backend/service"
)

// Config represents the configuration of the service
type Config struct {
	AppConfig                         config.AppConfig
	MongoConfig                       config.MongoConfig
	FirebaseConfig                    config.FirebaseConfig
	UserRepositoryConfig              repository.UserRepositoryConfig
	PatientRepositoryConfig           repository.PatientRepositoryConfig
	PatientAssessmentRepositoryConfig repository.PatientAssessmentRepositoryConfig
	NurseMonitoringRepositoryConfig   repository.NurseMonitoringRepositoryConfig
	MedicationRepositoryConfig        repository.MedicationRepositoryConfig
	NursingRepositoryConfig           repository.NursingRepositoryConfig
	HospitalRepositoryConfig          repository.HospitalRepositoryConfig
	WardRepositoryConfig              repository.WardRepositoryConfig
	RoomRepositoryConfig              repository.RoomRepositoryConfig
	CounterRepositoryConfig           repository.CounterRepositoryConfig
}

type mongoRepositories struct {
	userRepository              repository.UserRepository
	patientRepository           repository.PatientRepository
	patientAssessmentRepository repository.PatientAssessmentRepository
	nurseMonitoringRepository   repository.NurseMonitoringRepository
	medicationRepository        repository.MedicationRepository
	nursingRepository           repository.NursingRepository
	hospitalRepository          repository.HospitalRepository
	wardRepository              repository.WardRepository
	roomRepository              repository.RoomRepository
	counterRepository           repository.CounterRepository
}

// New injects the dependencies for the server
func New(c Config) {
	ctx := context.Background()

	e := echo.New()
	setupServer(ctx, e, c)

	authRepo, storageRepo, err := setupFirebase(ctx, c.FirebaseConfig)
	if err != nil {
		log.Panicf("error - [di.setupFirebase] unable to initialize Firebase client: %v", err)
	}

	mongoRepos := newMongoRepositories(ctx, c)

	service := service.New(service.Dependencies{
		UserRepository:              mongoRepos.userRepository,
		PatientRepository:           mongoRepos.patientRepository,
		PatientAssessmentRepository: mongoRepos.patientAssessmentRepository,
		NurseMonitoringRepository:   mongoRepos.nurseMonitoringRepository,
		MedicationRepository:        mongoRepos.medicationRepository,
		NursingRepository:           mongoRepos.nursingRepository,
		HospitalRepository:          mongoRepos.hospitalRepository,
		WardRepository:              mongoRepos.wardRepository,
		RoomRepository:              mongoRepos.roomRepository,
		AuthRepository:              authRepo,
		StorageRepository:           storageRepo,
		CounterRepository:           mongoRepos.counterRepository,
	})

	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.initServer] unable to start server: %v", err)
	}
}
