package domain

import (
	"context"
	"time"
)

type NotificationData struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"-"`
	Title     string    `json:"title"`
	Message   string    `json:"desc"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserUsecase represent the user's usecases
type NotificationUsecase interface {
	GetListNotification(ctx context.Context, userId int64) ([]*NotificationData, error)
	GetSingleNotification(ctx context.Context, id int64, userId int64) (*NotificationData, error)
}

// UserRepository represent the user's repository contract
type NotificationRepository interface {
	GetListNotification(ctx context.Context, userId int64) ([]*NotificationData, error)
	GetSingleNotification(ctx context.Context, id int64, userId int64) (*NotificationData, error)
	CreateNotification(ctx context.Context, title string, msg string, userId int64) error
}
