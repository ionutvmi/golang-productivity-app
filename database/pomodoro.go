package database

import (
	"app/models"
)

func PomodoroInsert(ps *models.Pomodoro) error {

	_, err := db.Exec(
		`INSERT INTO pomodoro (session_type, start_date, end_date) VALUES (?, ?, ?);`,
		ps.SessionType.ID,
		ps.StartDate,
		ps.EndDate,
	)

	if err != nil {
		return err
	}

	return nil
}

func PomodoroStats() *models.PomodoroStats {
	return &models.PomodoroStats{}
}
