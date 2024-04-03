package usecase

import (
	"L2/develop/dev11/internal/entity"
	"L2/develop/dev11/internal/repository"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type eventInteractor struct {
	repo repository.EventRepository
}

func NewEventInteractor(repo repository.EventRepository) *eventInteractor {
	return &eventInteractor{repo: repo}
}

func (i *eventInteractor) Create(ctx context.Context, event *entity.Event) error {
	err := i.repo.Create(ctx, event)
	if err != nil {
		return fmt.Errorf("error in eventInteractor.Create: %w", err)
	}

	return nil
}

func (i *eventInteractor) Update(ctx context.Context, event *entity.Event) error {
	err := i.repo.Update(ctx, event)
	if err != nil {
		return fmt.Errorf("error in eventInteractor.Update: %w", err)
	}

	return nil
}

func (i *eventInteractor) Delete(ctx context.Context, eventID uuid.UUID) error {
	err := i.repo.Delete(ctx, eventID)
	if err != nil {
		return fmt.Errorf("error in eventInteractor.Delete: %w", err)
	}

	return nil
}

func (i *eventInteractor) GetForDay(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	events, err := i.repo.GetForDay(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("error in eventInteractor.GetForDay: %w", err)
	}

	return events, nil
}

func (i *eventInteractor) GetForWeek(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	events, err := i.repo.GetForWeek(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("error in eventInteractor.GetForWeek: %w", err)
	}

	return events, nil
}

func (i *eventInteractor) GetForMonth(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	events, err := i.repo.GetForMonth(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("error in eventInteractor.GetForMonth: %w", err)
	}

	return events, nil
}
