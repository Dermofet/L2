// Package usecase provides interfaces for defining order interactor.
package usecase

import (
	"L2/develop/dev11/internal/entity"
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./interfaces.go -destination=usecases_mock.go -package=usecase

type EventInteractor interface {
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, event *entity.Event) error
	Delete(ctx context.Context, eventID uuid.UUID) error
	GetForDay(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
	GetForWeek(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
	GetForMonth(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error)
}
