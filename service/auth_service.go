package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/ShiftOver/shiftover-backend/pkg/utils"
	"github.com/pkg/errors"
)

func (s *service) SignUp(ctx context.Context, req dto.SignUpReq) (*dto.SignUpRes, error) {
	newUserID, err := s.counterRepository.GetCurrentUserIDCount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.SignUp]: unable to get current userId count")
	}
	newUserID++

	digits := utils.CountDigits(newUserID)
	userID := "N" + strings.Repeat("0", 7-digits) + strconv.Itoa(newUserID)

	err = s.counterRepository.IncrementUserIDCount(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.SignUp]: unable to increment userId count")
	}

	ward, err := s.wardRepository.Fetch(ctx, req.WardID)
	if err != nil {
		_ = s.counterRepository.DecrementUserIDCount(ctx)
		return nil, errors.Wrap(err, "error - [service.SignUp]: unable to fetch ward")
	}
	hospitalID := ward.HospitalID

	err = s.authRepository.SignUp(ctx, req, userID, hospitalID)
	if err != nil {
		_ = s.counterRepository.DecrementUserIDCount(ctx)
		return nil, errors.Wrap(err, "error - [service.SignUp]: unable to sign up user")
	}

	err = s.userRepository.Insert(ctx, dto.UserEntity{
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
		return nil, errors.Wrap(err, "error - [service.SignUp]: unable to insert user")
	}

	return &dto.SignUpRes{
		UserID:  userID,
		NurseID: req.NurseID,
		Email:   req.Email,
	}, nil
}
