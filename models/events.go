package models

import (
	"fmt"
	"time"

	"github.com/avijeetpandey/event-booking/db"
)

// Event model
type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)
	`

	preparedStatement, err := db.GlobalDB.Prepare(query)

	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	result, err := preparedStatement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * from events`
	rows, err := db.GlobalDB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := fmt.Sprintf("SELECT * FROM events WHERE id = %d", id)
	row := db.GlobalDB.QueryRow(query)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &event, nil
}
