// Package service provides the business logic service layer for the server
package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// GetUser fetches a user by their ID
func (s *service) GetUser(ctx context.Context, userID string) (*dto.UserEntity, error) {
	user, err := s.userRepository.Fetch(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
