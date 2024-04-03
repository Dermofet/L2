// Package db provides interfaces and methods for working with the database.

package db

import (
	"L2/develop/dev11/internal/entity"
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=source_mock.go -package=db

type EventSource interface {
	CreateEvent(ctx context.Context, event *entity.Event) error
	UpdateEvent(ctx context.Context, event *entity.Event) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
	GetEventForDay(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
	GetEventForWeek(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
	GetEventForMonth(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
}
