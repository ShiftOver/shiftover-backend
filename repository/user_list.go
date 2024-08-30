package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (r *userRepository) List(ctx context.Context) ([]*dto.UserEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	results := make([]*dto.UserEntity, 0)
	for cursor.Next(ctx) {
		result := new(dto.UserEntity)
		if err := cursor.Decode(result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
