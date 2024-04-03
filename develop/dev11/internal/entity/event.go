package entity

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID     uuid.UUID `json:"id,omitempty" db:"id"`
	Title  string    `json:"title,omitempty" db:"title"`
	Date   time.Time `json:"date,omitempty" db:"date"`
	UserID uuid.UUID `json:"user_id,omitempty" db:"user_id"`
}

func UnmarshalEvent(data []byte) (*Event, error) {
	u := &Event{}
	if err := json.Unmarshal(data, u); err != nil {
		return nil, err
	}
	return u, nil
}

func ParseFormEvent(form url.Values) (*Event, error) {
	id, err := uuid.Parse(form.Get("id"))
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", form.Get("date"))
	if err != nil {
		return nil, err
	}

	title := form.Get("title")
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}

	userID, err := uuid.Parse(form.Get("user_id"))
	if err != nil {
		return nil, err
	}

	return &Event{
		ID:     id,
		Title:  title,
		Date:   date,
		UserID: userID,
	}, nil
}

type Events []Event

func (e *Events) ToJSON() ([]byte, error) {
	data := map[string][]Event{"success": *e}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("can't marshal events: %v", err)
	}

	return jsonData, nil
}

func (e *Events) Add(event Event) {
	*e = append(*e, event)
}
