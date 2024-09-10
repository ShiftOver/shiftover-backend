package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/pkg/utils"
)

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func (s *service) SignUp(ctx context.Context, req *dto.SignUpReq) (*dto.SignUpRes, error) {
	newUserID, err := s.counterRepository.GetCurrentUserIDCount(ctx)
	if err != nil {
		return nil, err
	}
	newUserID++

	digits := countDigits(newUserID)
	userID := "N" + strings.Repeat("0", 7-digits) + strconv.Itoa(newUserID)

	err = s.counterRepository.IncrementUserIDCount(ctx)
	if err != nil {
		return nil, err
	}

	err = s.authRepository.SignUp(ctx, req, userID, "5")
	if err != nil {
		return nil, err
	}

	err = s.userRepository.InsertUser(ctx, &dto.UserEntity{
		UserID:            userID,
		NurseID:           req.NurseID,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		Email:             req.Email,
		ProfilePictureURL: req.ProfilePictureURL,
		WardID:            req.WardID,
		StartDate:         req.StartDate,
		DateOfBirth:       req.DateOfBirth,
		Position:          req.Position,
		Contact:           req.Contact,
		CreatedAt:         utils.LocalTime(),
		UpdatedAt:         utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}

	return &dto.SignUpRes{
		UserID:  userID,
		NurseID: req.NurseID,
		Email:   req.Email,
	}, nil
}
