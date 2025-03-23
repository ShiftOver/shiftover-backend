// Package service provides the business logic service layer for the server
package service

import (
	"context"
	"fmt"
	"time"

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

// InsertUser inserts a new user into the database
func (s *service) InsertUser(ctx context.Context, user dto.UserEntity) error {
	// Fetch the next user ID from the counter repository
	userID, err := s.counterRepository.GetCurrentUserIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertUser]: unable to fetch current user ID count")
	}

	// Increment the user ID counter
	err = s.counterRepository.IncrementUserIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertUser]: unable to increment user ID count")
	}

	// Set the generated UserID
	user.UserID = fmt.Sprintf("USER-%d", userID)

	// Set additional fields for the user
	user.CreatedAt = time.Now()

	// Insert the user into the database
	err = s.userRepository.Insert(ctx, user)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertUser]: unable to insert user")
	}

	return nil
}
