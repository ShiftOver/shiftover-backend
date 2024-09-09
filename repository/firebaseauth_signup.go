package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
)

// FirebaseAuthRepository represents the repository functions for the firebase auth
func (r *firebaseAuthRepository) SignUp(ctx context.Context, payload *dto.SignUpReq, userID, hospitalID string) error {
	params := (&auth.UserToCreate{}).
		UID(userID).
		Email(payload.Email).
		Password(payload.Password).
		DisplayName(fmt.Sprintf("%s %s", payload.FirstName, payload.LastName))

	user, err := r.authClient.CreateUser(ctx, params)
	if err != nil {
		return errors.Wrap(err, "error - [firebaseAuthRepository.SignUp]: unable to create user in firebase")
	}

	_ = user

	err = r.authClient.SetCustomUserClaims(ctx, user.UID, map[string]interface{}{
		"uid":         userID,
		"nurse_id":    payload.NurseID,
		"ward_id":     payload.WardID,
		"hospital_id": hospitalID,
	})
	if err != nil {
		return errors.Wrap(err, "error - [firebaseAuthRepository.SignUp]: unable to set custom claims for user")
	}

	return nil
}
