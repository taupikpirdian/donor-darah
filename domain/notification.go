package domain

import (
	"context"
	"time"
)

type NotificationData struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserUsecase represent the user's usecases
type NotificationUsecase interface {
	GetListNotification(ctx context.Context) ([]*NotificationData, error)
	GetSingleNotification(ctx context.Context, id int64) (*NotificationData, error)
}

// UserRepository represent the user's repository contract
type NotificationRepository interface {
	GetListNotification(ctx context.Context) ([]*NotificationData, error)
	GetSingleNotification(ctx context.Context, id int64) (*NotificationData, error)
}
