package domain

import (
	"context"
	"time"
)

// User is representing the User data struct
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Register(ctx context.Context, us *User) (User, error)
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Register(ctx context.Context, us *User) (User, error)
}
