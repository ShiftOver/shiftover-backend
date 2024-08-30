package di

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/ShiftOver/shiftover-backend/config"
	"github.com/ShiftOver/shiftover-backend/repository"
	"github.com/labstack/gommon/log"
)

func setupFirebase(pctx context.Context, c config.FirebaseConfig) (repository.FirebaseAuthRepository, repository.FirebaseStorageRepository, error) {
	firebaseConfig := &firebase.Config{
		ProjectID:     c.ProjectID,
		StorageBucket: c.StorageBucket,
	}

	var fb *firebase.App
	var err error

	if c.CredentialsJSON == "" {
		fb, err = firebase.NewApp(pctx, firebaseConfig)
		if err != nil {
			return nil, nil, err
		}
	} else {
		fb, err = firebase.NewApp(pctx, firebaseConfig, option.WithCredentialsJSON([]byte(c.CredentialsJSON)))
		if err != nil {
			return nil, nil, err
		}
	}

	authClient, err := fb.Auth(pctx)
	if err != nil {
		return nil, nil, err
	}

	authRepo := repository.NewFirebaseAuthRepository(repository.AuthDependencies{
		AuthClient: authClient,
	})

	log.Infof("[di.setupFirebase] firebase authentication initialized")

	storageClient, err := fb.Storage(pctx)
	if err != nil {
		return nil, nil, err
	}

	storageRepo := repository.NewFirebaseStorageRepository(repository.StorageDependencies{
		StorageClient:         storageClient,
		Bucket:                c.StorageBucket,
		UserProfileObjectPath: c.UserProfileObjectPath,
	})

	log.Infof("[di.setupFirebase] firebase storage initialized")

	return authRepo, storageRepo, nil
}
