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
	userRepository repository.UserRepository
}

// Dependencies represents the dependencies for the service
type Dependencies struct {
	UserRepository repository.UserRepository
}

// New creates a new service
func New(d Dependencies) Port {
	return &service{
		userRepository: d.UserRepository,
	}
}
