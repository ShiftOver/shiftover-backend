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
