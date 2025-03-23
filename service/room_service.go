package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/ShiftOver/shiftover-backend/dto"
)

func (s *service) GetRoom(ctx context.Context, roomID string) (*dto.GetRoomResponse, error) {
	room, err := s.roomRepository.Fetch(ctx, roomID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetRoom]: unable to fetch room")
	}

	if room.CurrentPatientID == "" {
		return &dto.GetRoomResponse{
			RoomID:    room.RoomID,
			RoomName:  room.RoomName,
			WardID:    room.WardID,
			UpdatedAt: room.UpdatedAt,
			CreatedAt: room.CreatedAt,
		}, nil
	}

	currentPatient, err := s.patientRepository.Fetch(ctx, room.CurrentPatientID)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.GetRoom]: unable to fetch current patient")
	}

	return &dto.GetRoomResponse{
		RoomID:         room.RoomID,
		RoomName:       room.RoomName,
		WardID:         room.WardID,
		CurrentPatient: *currentPatient,
		UpdatedAt:      room.UpdatedAt,
		CreatedAt:      room.CreatedAt,
	}, nil
}

func (s *service) ListRooms(ctx context.Context) ([]*dto.GetRoomResponse, error) {
	rooms, err := s.roomRepository.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error - [service.ListRooms]: unable to list rooms")
	}

	results := make([]*dto.GetRoomResponse, 0)
	for _, room := range rooms {
		if room.CurrentPatientID == "" {
			results = append(results, &dto.GetRoomResponse{
				RoomID:    room.RoomID,
				RoomName:  room.RoomName,
				WardID:    room.WardID,
				UpdatedAt: room.UpdatedAt,
				CreatedAt: room.CreatedAt,
			})
			continue
		}

		currentPatient, err := s.patientRepository.Fetch(ctx, room.CurrentPatientID)
		if err != nil {
			return nil, errors.Wrap(err, "error - [service.ListRooms]: unable to fetch current patient")
		}

		results = append(results, &dto.GetRoomResponse{
			RoomID:         room.RoomID,
			RoomName:       room.RoomName,
			WardID:         room.WardID,
			CurrentPatient: *currentPatient,
			UpdatedAt:      room.UpdatedAt,
			CreatedAt:      room.CreatedAt,
		})
	}

	return results, nil
}

func (s *service) InsertRoom(ctx context.Context, room dto.RoomEntity) error {
	// Fetch the next room ID from the counter repository
	roomID, err := s.counterRepository.GetCurrentRoomIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertRoom]: unable to fetch current room ID count")
	}

	// Increment the room ID counter
	err = s.counterRepository.IncrementRoomIDCount(ctx)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertRoom]: unable to increment room ID count")
	}

	// Set the room ID
	room.RoomID = fmt.Sprintf("ROOM-%d", roomID)

	// Insert the room into the database
	err = s.roomRepository.Insert(ctx, room)
	if err != nil {
		return errors.Wrap(err, "error - [service.InsertRoom]: unable to insert room")
	}

	return nil
}
