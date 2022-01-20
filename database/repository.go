package database

// Repository represent the repositories
type Repository interface {
	Close()
	FindByID(id string) (e EventModel, err error)
	Find() (es []EventModel, err error)
	Create(e EventModel) (err error)
	Update(e EventModel) (err error)
	Delete(id string) (err error)
	CreateTable() (err error)
}

type EventModel struct {
	Id   int
	Name string
	Year int
}
