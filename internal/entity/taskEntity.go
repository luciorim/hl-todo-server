package entity

import "time"

type Task struct {
	ID       string
	Title    string
	ActiveAt time.Time
	Status   bool
}
