// Package repository provides interfaces for interacting with order repository.
package repository

import (
	"L2/develop/dev11/internal/entity"
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./interfaces.go -destination=repositories_mock.go -package=repository

type EventRepository interface {
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, eventID uuid.UUID) error
	GetForDay(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
	GetForWeek(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
	GetForMonth(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
}
