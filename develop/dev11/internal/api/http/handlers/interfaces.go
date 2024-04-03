// Package handlers provides interfaces for order handlers.
package handlers

import (
	"net/http"
)

//go:generate mockgen -source=interfaces.go -destination=handlers_mock.go -package=handlers

type EventHandlers interface {
	CreateHandler(http.ResponseWriter, *http.Request)
	UpdateHandler(http.ResponseWriter, *http.Request)
	DeleteHandler(http.ResponseWriter, *http.Request)
	GetForDayHandler(http.ResponseWriter, *http.Request)
	GetForWeekHandler(http.ResponseWriter, *http.Request)
	GetForMonthHandler(http.ResponseWriter, *http.Request)
}
