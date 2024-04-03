package db

import (
	"L2/develop/dev11/internal/entity"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s *source) CreateEvent(ctx context.Context, event *entity.Event) error {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	row := s.db.QueryRowContext(
		dbCtx,
		"INSERT INTO events (id, title, date, user_id) VALUES ($1, $2, $3, $4);",
		event.ID, event.Title, event.Date, event.UserID,
	)
	if err := row.Err(); err != nil {
		return fmt.Errorf("can't exec query: %v", err)
	}

	return nil
}

func (s *source) UpdateEvent(ctx context.Context, event *entity.Event) error {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	row := s.db.QueryRowContext(
		dbCtx,
		"UPDATE events SET title = $1, date = $2 WHERE id = $3;",
		event.Title, event.Date, event.ID,
	)
	if err := row.Err(); err != nil {
		return fmt.Errorf("can't exec query: %v", err)
	}

	return nil
}

func (s *source) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	row := s.db.QueryRowContext(
		dbCtx,
		"DELETE events WHERE id = $1",
		eventID,
	)
	if err := row.Err(); err != nil {
		return fmt.Errorf("can't exec query: %v", err)
	}

	return nil
}

func (s *source) GetEventForDay(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	// Вычисляем начало и конец указанного дня
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	rows, err := s.db.QueryxContext(
		dbCtx,
		"SELECT * FROM events WHERE user_id = $1 AND date >= $2 AND date < $3",
		userID, startOfDay, endOfDay,
	)
	if err != nil {
		return nil, fmt.Errorf("can't exec query: %v", err)
	}
	defer rows.Close()

	events := &entity.Events{}

	for rows.Next() {
		var event entity.Event
		if err := rows.StructScan(&event); err != nil {
			return nil, fmt.Errorf("can't scan event: %v", err)
		}
		events.Add(event)
	}

	return events, nil
}

func (s *source) GetEventForWeek(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	// Вычисляем начало и конец недели
	startOfWeek := date
	endOfWeek := date.AddDate(0, 0, 7)

	rows, err := s.db.QueryxContext(
		dbCtx,
		"SELECT * FROM events WHERE user_id = $1 AND date >= $2 AND date <= $3",
		userID, startOfWeek, endOfWeek,
	)
	if err != nil {
		return nil, fmt.Errorf("can't exec query: %v", err)
	}
	defer rows.Close()

	events := &entity.Events{}

	for rows.Next() {
		var event entity.Event
		if err := rows.StructScan(&event); err != nil {
			return nil, fmt.Errorf("can't scan event: %v", err)
		}
		events.Add(event)
	}

	return events, nil
}

func (s *source) GetEventForMonth(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Events, error) {
	dbCtx, dbCancel := context.WithTimeout(ctx, QueryTimeout)
	defer dbCancel()

	// Вычисляем начало и конец месяца
	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	// Выполняем запрос к базе данных
	rows, err := s.db.QueryxContext(
		dbCtx,
		"SELECT * FROM events WHERE user_id = $1 AND date >= $2 AND date <= $3",
		userID, startOfMonth, endOfMonth,
	)
	if err != nil {
		return nil, fmt.Errorf("can't exec query: %v", err)
	}
	defer rows.Close()

	events := &entity.Events{}

	// Обрабатываем результаты запроса
	for rows.Next() {
		var event entity.Event
		if err := rows.StructScan(&event); err != nil {
			return nil, fmt.Errorf("can't scan event: %v", err)
		}
		events.Add(event)
	}

	return events, nil
}
