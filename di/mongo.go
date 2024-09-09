package di

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ShiftOver/shiftover-backend/repository"
)

func newMongoRepositories(pctx context.Context, c Config) (repository.UserRepository, repository.PatientRepository, repository.HospitalRepository, repository.WardRepository, repository.RoomRepository) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", url.QueryEscape(c.MongoConfig.Username), url.QueryEscape(c.MongoConfig.Password), c.MongoConfig.HostName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Panicf("error - [di.newMongoRepositories] cannot connect to database: %v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Panicf("error - [di.newMongoRepositories] error pinging to database: %v", err)
	}

	log.Info("[di.newMongoRepositories] connected to database")

	userColl := client.Database(c.MongoConfig.DBName).Collection(c.UserRepositoryConfig.CollectionName)
	userRepo := repository.NewUserRepository(*userColl)

	patientColl := client.Database(c.MongoConfig.DBName).Collection(c.PatientRepositoryConfig.CollectionName)
	patientRepo := repository.NewPatientRepository(*patientColl)

	hospitalColl := client.Database(c.MongoConfig.DBName).Collection(c.HospitalRepositoryConfig.CollectionName)
	hospitalRepo := repository.NewHospitalRepository(*hospitalColl)

	wardColl := client.Database(c.MongoConfig.DBName).Collection(c.WardRepositoryConfig.CollectionName)
	wardRepo := repository.NewWardRepository(*wardColl)

	roomColl := client.Database(c.MongoConfig.DBName).Collection(c.RoomRepositoryConfig.CollectionName)
	roomRepo := repository.NewRoomRepository(*roomColl)

	return userRepo, patientRepo, hospitalRepo, wardRepo, roomRepo
}
