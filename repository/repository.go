// Package repository provides the repository interfaces for the server
package repository

//go:generate mockgen -destination=./mock/mock_repository.go -package=mock -source=./repository.go UserRepository

// UserRepository represents the repository functions for the user collection
type UserRepository interface {
}
