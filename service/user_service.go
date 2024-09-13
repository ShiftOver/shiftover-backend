// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// GetUser fetches a user by their ID
func (s *service) GetUser(ctx context.Context, userID string) (*dto.UserEntity, error) {
	user, err := s.userRepository.Fetch(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetUser]: unable to fetch user")
	}
	return user, nil
}
