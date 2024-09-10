package repository

import (
	"firebase.google.com/go/auth"
)

type firebaseAuthRepository struct {
	authClient *auth.Client
}

// AuthDependencies represents the dependencies for the firebase auth repository
type AuthDependencies struct {
	AuthClient *auth.Client
}

// NewFirebaseAuthRepository creates a new firebase auth repository
func NewFirebaseAuthRepository(d AuthDependencies) FirebaseAuthRepository {
	return &firebaseAuthRepository{
		authClient: d.AuthClient,
	}
}
