package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/pkg/errors"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// SignUp creates a new user in firebase
func (r *firebaseAuthRepository) SignUp(ctx context.Context, payload dto.SignUpReq, userID, hospitalID string) error {
	params := (&auth.UserToCreate{}).
		UID(userID).
		Email(payload.Email).
		Password(payload.Password).
		DisplayName(fmt.Sprintf("%s %s", payload.FirstName, payload.LastName))

	_, err := r.authClient.CreateUser(ctx, params)
	if err != nil {
		return errors.Wrap(err, "error - [firebaseAuthRepository.SignUp]: unable to create user in firebase")
	}

	err = r.authClient.SetCustomUserClaims(ctx, userID, map[string]interface{}{
		"nurseId":    payload.NurseID,
		"wardId":     payload.WardID,
		"hospitalId": hospitalID,
	})
	if err != nil {
		return errors.Wrap(err, "error - [firebaseAuthRepository.SignUp]: unable to set custom claims for user")
	}

	return nil
}
