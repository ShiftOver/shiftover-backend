package repository

import (
	"firebase.google.com/go/storage"
)

type firebaseStorageRepository struct {
	storageClient         *storage.Client
	bucket                string
	userProfileObjectPath string
}

// StorageDependencies represents the dependencies for the firebase storage repository
type StorageDependencies struct {
	StorageClient         *storage.Client
	Bucket                string
	UserProfileObjectPath string
}

// NewFirebaseStorageRepository creates a new firebase storage repository
func NewFirebaseStorageRepository(d StorageDependencies) FirebaseStorageRepository {
	return &firebaseStorageRepository{
		storageClient:         d.StorageClient,
		bucket:                d.Bucket,
		userProfileObjectPath: d.UserProfileObjectPath,
	}
}
