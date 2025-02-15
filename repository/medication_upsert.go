package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Insert inserts a new medication into the database
func (r *medicationRepository) Upsert(ctx context.Context, filter map[string]interface{}, update map[string]interface{}) error {
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, bson.M{"$set": update}, opts)
	if err != nil {
		return errors.Wrap(err, "error - [medicationRepository.Upsert]: unable to upsert medication")
	}
	return nil
}
