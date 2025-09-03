package models

import (
	"fmt"
	"time"

	"antero.com/event_booking/db"
)

type Event struct {
	ID          int64     `json:"id"`
	EventName   string    `json:"eventName"   binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"    binding:"required"`
	Capacity    int       `json:"capacity"`
	DateTime    time.Time `json:"dateTime"    binding:"required"`
	UserID      int64     `json:"userId"`
}

var events = []Event{}

func (e *Event) Save() error {
	// later: add it to a database
	query := `INSERT INTO events(eventName,description,location,capacity, dateTime, user_id)
	VALUES (?,?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare insert: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.EventName, e.Description, e.Location, e.Capacity, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("stmt.Exec: %w", err)
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.EventName, &event.Description, &event.Location, &event.Capacity, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventBYId(id int64) (*Event, error) {
	query := "SELECT * from events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.EventName, &event.Description, &event.Location, &event.Capacity, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Update updates the current Event record in the database with the values from the Event struct.
// It returns an error if the update operation fails.
func (event Event) Update() error {
	query := `
		UPDATE events
		SET eventName = ?, description = ?, location = ?, capacity = ?, dateTime = ?
		WHERE id = ?
		`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.EventName, event.Description, event.Location, event.Capacity, event.DateTime, event.ID)
	return err

}

// Delete removes the current Event instance from the data store.
// It returns an error if the deletion fails.
func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return nil

}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return nil

}
