// Package http provides functionality for handling HTTP requests.
package http

import (
	"fmt"
	"net/http"

	"L2/develop/dev11/internal/api/http/handlers"
	"L2/develop/dev11/internal/api/http/middleware"
	"L2/develop/dev11/internal/db"
	"L2/develop/dev11/internal/repository"
	"L2/develop/dev11/internal/usecase"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// routerHandlers contains handlers for router.
type routerHandlers struct {
	eventHandlers handlers.EventHandlers
}

// router represents an HTTP router.
type router struct {
	mux      http.Handler
	db       *sqlx.DB
	handlers routerHandlers
	logger   *zap.Logger
}

// NewRouter creates a new instance of HTTP router.
func NewRouter(db *sqlx.DB, logger *zap.Logger) *router {
	return &router{
		mux:    http.NewServeMux(),
		db:     db,
		logger: logger,
	}
}

// Init initializes the HTTP router.
func (r *router) Init() error {
	err := r.registerRoutes()
	if err != nil {
		return fmt.Errorf("can't init router: %w", err)
	}

	return nil
}

// registerRoutes registers routes in the HTTP router.
func (r *router) registerRoutes() error {
	mux := &http.ServeMux{}
	handler := middleware.Recovery(mux)
	handler = middleware.Logging(handler)

	pgSource := db.NewSource(r.db)
	eventRepository := repository.NewEventRepository(pgSource)
	eventInteractor := usecase.NewEventInteractor(eventRepository)
	r.handlers.eventHandlers = handlers.NewEventHandlers(eventInteractor)

	mux.HandleFunc("/create_event", r.handlers.eventHandlers.CreateHandler)
	mux.HandleFunc("/update_event", r.handlers.eventHandlers.UpdateHandler)
	mux.HandleFunc("/delete_event", r.handlers.eventHandlers.DeleteHandler)
	mux.HandleFunc("/events_for_day", r.handlers.eventHandlers.GetForDayHandler)
	mux.HandleFunc("/events_for_week", r.handlers.eventHandlers.GetForWeekHandler)
	mux.HandleFunc("/events_for_month", r.handlers.eventHandlers.GetForMonthHandler)

	r.mux = handler

	return nil
}
