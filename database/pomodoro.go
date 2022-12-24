package database

import (
	"app/models"
	"log"
	"time"
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
	var stats = &models.PomodoroStats{
		Today: 0,
		Week:  0,
		Month: 0,
		Year:  0,
	}
	var now = time.Now().UTC()

	var startOfDay = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	var err error

	err = db.Get(&stats.Today, `SELECT COUNT(session_id) 
	FROM  pomodoro
		WHERE start_date >= ?
		`, startOfDay)

	if err != nil {
		log.Println("Failed to calculate stats for today: ", err.Error())
		stats.Today = 0
	}

	var startOfWeek = getStartDayOfWeek(now)
	err = db.Get(&stats.Week, `SELECT COUNT(session_id)
		FROM pomodoro
		WHERE start_date >= ?
	`, startOfWeek)

	if err != nil {
		log.Println("Failed to calculate stats for this week: ", err.Error())
		stats.Week = 0
	}

	var startOfMonth = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	err = db.Get(&stats.Week, `SELECT COUNT(session_id)
		FROM  pomodoro
		WHERE start_date >= ?
	`, startOfMonth)

	if err != nil {
		log.Println("Failed to calculate stats for this month: ", err.Error())
		stats.Month = 0
	}

	var startOfYear = time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	err = db.Get(&stats.Week, `SELECT COUNT(session_id)
		FROM  pomodoro
		WHERE start_date >= ?
	`, startOfYear)

	if err != nil {
		log.Println("Failed to calculate stats for this year: ", err.Error())
		stats.Year = 0
	}

	return stats
}

func getStartDayOfWeek(tm time.Time) time.Time {
	weekday := time.Duration(tm.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := tm.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, tm.Location())
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour)
}
