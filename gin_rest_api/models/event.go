package models

import (
	"rest-api/db"
	"time"
)

/* Make sure to name anything with first character to be upperCase */

// Shape of an Event

// To make variable of struct to be required in request object - use binding
type Event struct {
	ID int64
	Title string 		`binding:"required"`
	Description string  `binding:"required"`
	Location string     `binding:"required"`
	DateTime time.Time  `binding:"required"`
	UserID int64
}

// var events = []Event{}

// Adding method to event struct

func (e *Event) Save() error {
	// add to db

	// query : TO AVOID SQL INJECTION ATTACKS WE USE QUESTION marks
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}
	// To close this exec statement we can use defer keyword
	defer stmt.Close()

	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	//fetch from db

	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	// Loop over the rows and scan the variables and append in events slice
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err!= nil {
            return nil, err
        }
		events = append(events, event)
	}

    return events, nil
}

func GetEventById(id int64) (*Event, error) {
	// Use Pointer to event in return so that we can handle nil return
	query := `SELECT * FROM events WHERE id = ?`

	// QueryRow give single row
    row := db.DB.QueryRow(query, id)

    var event Event
    err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	// update in db

    query := `UPDATE events SET name =?, description =?, location =?, dateTime =? WHERE id =?`

    stmt, err := db.DB.Prepare(query)
    
    if err!= nil {
        return err
    }
    // To close this exec statement we can use defer keyword
    defer stmt.Close()

    _, err = stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.ID)

    return err
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id =?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}