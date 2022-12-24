package models

import "time"

type PomodoroType struct {
	ID   int `db:"type_id"`
	Name int `db:"type_name"`
}

type Pomodoro struct {
	SessionId   int       `db:"session_id"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
	SessionType PomodoroType
}

type PomodoroStats struct {
	Today int
	Week  int
	Month int
	Year  int
}
