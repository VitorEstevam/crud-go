package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	db *sql.DB
}

func NewRepository(dialect, dsn string, idleConn, maxConn int) (Repository, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}
func (r *repository) CreateTable() (err error) {
	query := `
    CREATE TABLE IF NOT EXISTS event(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        year INTEGER NOT NULL
    );
    `
	_, err = r.db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return
}

// func DropTable(db *sql.DB) error {
// 	query := `DROP TABLE IF EXISTS  EventModel;`

// 	_, err := db.Exec(query)
// 	return err
// }

func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) FindByID(id string) (e EventModel, err error) {
	query := fmt.Sprintf("SELECT * FROM event WHERE id = %v", id)
	rows, err := r.db.Query(query)
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

func (r *repository) Find() (es []EventModel, err error) {
	query := "SELECT * FROM event"

	rows, err := r.db.Query(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	for rows.Next() {
		EventModel := EventModel{}

		err = rows.Scan(&EventModel.Id, &EventModel.Name, &EventModel.Year)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		es = append(es, EventModel)
	}
	return
}

func (r *repository) Create(event EventModel) (err error) {
	query := fmt.Sprintf("INSERT INTO event (name, year) VALUES (\"%v\", %v)", event.Name, event.Year)

	_, err = r.db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	return
}

func (r *repository) Update(event EventModel) (err error) {
	query := fmt.Sprintf("UPDATE event SET name = \"%v\", year = %v WHERE id = %v", event.Name, event.Year, event.Id)

	_, err = r.db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return
}

func (r *repository) Delete(id string) (err error) {
	query := fmt.Sprintf("DELETE FROM event WHERE id = %v", id)
	_, err = r.db.Exec(query)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	return
}
