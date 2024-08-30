package service

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// ListUser lists all the users
func (s *service) ListUser(ctx context.Context) ([]*dto.UserEntity, error) {
	return s.userRepository.List(ctx)
}
