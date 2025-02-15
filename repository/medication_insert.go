package repository

import (
	"context"

	"github.com/ShiftOver/shiftover-backend/dto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Insert inserts a new medication into the database
func (r *medicationRepository) Insert(ctx context.Context, medication dto.MedicationEntity) error {
	_, err := r.collection.InsertOne(ctx, medication, options.InsertOne())
	if err != nil {
		return errors.Wrap(err, "error - [medicationRepository.Insert]: unable to insert medication")
	}
	return nil
}
