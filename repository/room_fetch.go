package repository

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ShiftOver/shiftover-backend/dto"
)

// Fetch fetches a room by room ID
func (r *roomRepository) Fetch(ctx context.Context, roomID string) (*dto.RoomEntity, error) {
	room := r.collection.FindOne(ctx, bson.D{primitive.E{Key: "roomId", Value: roomID}})

	var entity dto.RoomEntity
	if err := room.Decode(&entity); err != nil {
		return nil, errors.Wrap(err, "error - [roomRepository.Fetch]: unable to decode result")
	}

	return &entity, nil
}

// List fetches all rooms in the database
func (r *roomRepository) List(ctx context.Context) ([]*dto.RoomEntity, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "error - [roomRepository.List]: unable to list rooms")
	}

	results := make([]*dto.RoomEntity, 0)
	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "error - [roomRepository.List]: unable to decode result")
	}

	return results, nil
}
