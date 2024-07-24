package dto

import (
	"errors"
	"time"
)

type TaskRequestDto struct {
	Title    string `json:"title"`
	ActiveAt string `json:"active_at"`
}

type TaskResponseDto struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"active_at"`
	Done     bool      `json:"is_done"`
}

func (t *TaskRequestDto) ValidateRequest() error {

	// Validate title
	if len(t.Title) > 200 {
		return errors.New("title should be less than 200 characters")
	}

	// Validate date
	const dateFormat = "2006-01-02" // Correct date format
	if _, err := time.Parse(dateFormat, t.ActiveAt); err != nil {
		return errors.New("active_at should have this format: yyyy-MM-dd")
	}

	return nil
}
