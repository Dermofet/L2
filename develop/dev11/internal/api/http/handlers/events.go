package handlers

import (
	"L2/develop/dev11/internal/entity"
	"L2/develop/dev11/internal/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type eventHandlers struct {
	interactor usecase.EventInteractor
}

func NewEventHandlers(interactor usecase.EventInteractor) *eventHandlers {
	return &eventHandlers{
		interactor: interactor,
	}
}

func (h *eventHandlers) CreateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Can't parse form", http.StatusBadRequest)
		return
	}

	event, err := entity.ParseFormEvent(req.Form)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.interactor.Create(req.Context(), event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *eventHandlers) UpdateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	err := req.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	event, err := entity.ParseFormEvent(req.Form)
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.interactor.Update(req.Context(), event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *eventHandlers) DeleteHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(req.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse user_id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.interactor.Delete(req.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *eventHandlers) GetForDayHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(req.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse user_id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", req.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse date: %s", err.Error()), http.StatusBadRequest)
		return
	}

	events, err := h.interactor.GetForDay(req.Context(), userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := events.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(body)
}

func (h *eventHandlers) GetForWeekHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(req.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse user_id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", req.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse date: %s", err.Error()), http.StatusBadRequest)
		return
	}

	events, err := h.interactor.GetForWeek(req.Context(), userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := events.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(body)
}

func (h *eventHandlers) GetForMonthHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(req.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse user_id: %s", err.Error()), http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", req.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Can't parse date: %s", err.Error()), http.StatusBadRequest)
		return
	}

	events, err := h.interactor.GetForMonth(req.Context(), userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := events.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
