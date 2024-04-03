package repository

import (
	"L2/develop/dev11/internal/db"
	"L2/develop/dev11/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type eventRepository struct {
	source db.EventSource
}

func NewEventRepository(source db.EventSource) *eventRepository {
	return &eventRepository{
		source: source,
	}
}

func (r *eventRepository) Create(ctx context.Context, event *entity.Event) error {
	err := r.source.CreateEvent(ctx, event)
	if err != nil {
		return fmt.Errorf("error in eventRepository.Create: %w", err)
	}

	return nil
}

func (r *eventRepository) Update(ctx context.Context, event *entity.Event) error {
	err := r.source.UpdateEvent(ctx, event)
	if err != nil {
		return fmt.Errorf("error in eventRepository.Update: %w", err)
	}

	return nil
}

func (r *eventRepository) Delete(ctx context.Context, eventID uuid.UUID) error {
	err := r.source.DeleteEvent(ctx, eventID)
	if err != nil {
		return fmt.Errorf("error in eventRepository.Delete: %w", err)
	}

	return nil
}

func (r *eventRepository) GetForDay(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	events, err := r.source.GetEventForDay(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("error in eventRepository.GetForDay: %w", err)
	}

	return events, nil
}

func (r *eventRepository) GetForWeek(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	events, err := r.source.GetEventForWeek(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("error in eventRepository.GetForWeek: %w", err)
	}

	return events, nil
}

func (r *eventRepository) GetForMonth(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	events, err := r.source.GetEventForMonth(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("error in eventRepository.GetForMonth: %w", err)
	}

	return events, nil
}
