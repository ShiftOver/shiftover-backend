package repository

import (
	"firebase.google.com/go/auth"
)

type firebaseAuthRepository struct {
	authClient *auth.Client
}

// AuthDependencies represents the dependencies for the repository
type AuthDependencies struct {
	AuthClient *auth.Client
}

// NewFirebaseAuthRepository creates a new repository
func NewFirebaseAuthRepository(d AuthDependencies) FirebaseAuthRepository {
	return &firebaseAuthRepository{
		authClient: d.AuthClient,
	}
}
