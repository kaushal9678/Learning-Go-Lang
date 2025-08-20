package models

import (
	"time"

	"example.com/rest-api-go/db"
)

type Event struct{
	ID	int64 `json:"id"`
	Name string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
	DateTime time.Time `json:"date" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID int `json:"user_id"`
}

var events = []Event{}
func (e Event) Save() error{
	events = append(events, e)
	query := `INSERT INTO events (name, location, date, description, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if(err != nil) {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Location, e.DateTime, e.Description, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id
	return nil
}
func GetAllEvents() []Event {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.UserID); err != nil {
			panic(err)
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	if len(events) == 0 {
		return nil
	}
	// Return a copy of the events slice to avoid external modification
	eventsCopy := make([]Event, len(events))
	copy(eventsCopy, events)	
	return events
}
func GetEventById(id int64)(*Event, error){
	query := `SELECT * FROM events WHERE id = ?`;
	row := db.DB.QueryRow(query,id);

	var event Event;
	err := row.Scan(&event.ID, &event.Name, &event.Location, &event.DateTime, &event.Description, &event.UserID);
	if err != nil {
		return nil, err
	}
	return &event, nil
}
func (event Event) Update() error {
	query := `UPDATE events SET name = ?, location = ?, date = ?, description = ?, user_id = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Location, event.DateTime, event.Description, event.UserID, event.ID)
	return err
}
func (event Event)DeleteEvent() error{
	query := `DELETE FROM events WHERE id = ?`;
	stmt, err := db.DB.Prepare(query);
	if err != nil{
		return err;
	}
	defer stmt.Close();
	_, err = stmt.Exec(event.ID)
	if err != nil{
		return err
	}
	return nil
}