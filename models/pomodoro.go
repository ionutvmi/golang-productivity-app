package models

import "time"

type PomodoroSession struct {
	SessionId   int       `db:"session_id"`
	CreatedDate time.Time `db:"created_date"`
}
