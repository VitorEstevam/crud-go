package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func init() {

}

type Event struct {
	Id   int
	Name string
	Year int
}

func Connect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DropTable(db *sql.DB) error {
	query := `DROP TABLE IF EXISTS  event;`

	_, err := db.Exec(query)
	return err
}

func CreateTable(db *sql.DB) error {
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

func InsertEvent(db *sql.DB, e Event) (res sql.Result, err error) {
	query := fmt.Sprintf("INSERT INTO event (name, year) VALUES (\"%v\", %v)", e.Name, e.Year)

	res, err = db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	return
}

func GetEvent(db *sql.DB, id int) (e Event, err error) {
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

func GetEvents(db *sql.DB) (events []Event, err error) {
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

func UpdateEvent(db *sql.DB, e Event) (res sql.Result, err error) {
	query := fmt.Sprintf("UPDATE event SET name = \"%v\", year = %v WHERE id = %v", e.Name, e.Year, e.Id)

	res, err = db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return res, err
}

func DeleteEvent(db *sql.DB, id int) (res sql.Result, err error) {
	query := fmt.Sprintf("DELETE FROM event WHERE id = %v", id)
	res, err = db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	return
}
