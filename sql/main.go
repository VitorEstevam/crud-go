package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Event struct {
	Id   int
	Name string
	Year int
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func dropTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS  event;`

	_, err := db.Exec(query)
	return err
}

func createTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS event(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        year INTEGER NOT NULL
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return err
}

func insertEvent(db *sql.DB, e Event) (res sql.Result, err error) {
	query := fmt.Sprintf("INSERT INTO event (name, year) VALUES (\"%v\", %v)", e.Name, e.Year)

	res, err = db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	return
}

func getEvent(db *sql.DB, id int) (e Event, err error) {
	query := fmt.Sprintf("SELECT * FROM event WHERE id = %v", id)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	rows.Next()
	err = rows.Scan(&e.Id, &e.Name, &e.Year)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return
}

func getEvents(db *sql.DB) (events []Event, err error) {
	query := "SELECT * FROM event"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	for rows.Next() {
		event := Event{}

		err = rows.Scan(&event.Id, &event.Name, &event.Year)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		events = append(events, event)
	}
	return
}

func updateEvent(db *sql.DB, e Event) (res sql.Result, err error) {
	query := fmt.Sprintf("UPDATE event SET name = \"%v\", year = %v WHERE id = %v", e.Name, e.Year, e.Id)

	res, err = db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return res, err
}

func deleteEvent(db *sql.DB, id int) (res sql.Result, err error) {
	query := fmt.Sprintf("DELETE FROM event WHERE id = %v", id)
	res, err = db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	return
}

func main() {
	// db := dbConn()
	// dropTable(db)
	// createTable(db)
	// e := Event{Name: "cccccccc", Year: 2021}
	// res, err := insertEvent(db, e)
	// fmt.Println(res, err)

	// res2, _ := deleteEvent(db, 3)
	// fmt.Println(res2)
	// db.Close()
}
