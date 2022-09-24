package factory

import (
	"agmc-day-6/database"
	"agmc-day-6/internal/repository"
)

type Factory struct {
	BookRepository repository.Book
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewBook(db),
		repository.NewUser(db),
	}
}
