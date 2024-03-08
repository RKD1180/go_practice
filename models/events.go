package models

import (
	"greeter/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() error {
	queery := `INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(queery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	queery := "SELECT * FROM events"
	rows, err := db.DB.Query(queery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		rows.Scan(&event.Id, &event.Name, &event.DateTime, &event.Description, &event.Location, &event.UserId)
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	queery := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(queery, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, date_time = ?
		WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return err
}

func (event Event) Delete() error {
	query := `
	DELETE events 
	WHERE id = ?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Id)
	return err
}
