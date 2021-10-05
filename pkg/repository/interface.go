package repository

import "context"

// VideoRepo provides video data
type VideoRepo interface {
	GetAll(ctx context.Context) error
}

// UserRepo provides users data
type UserRepo interface {
	GetAll(ctx context.Context) error
}
