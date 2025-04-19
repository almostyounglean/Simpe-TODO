package model

import "time"

type Task struct {
	Id          int
	Name        string
	Meta        string
	IsCompleted bool
	CreatedAt   time.Time
}
