package repositories

import "github.com/dibimbing-satkom-indo/onion-architecture-go/entities"

type UserRepository struct {
	db any
}

type UserRepositoryInterface interface {
	GetByID(id int) []entities.User
}

func (repo UserRepository) GetByID(id int) []entities.User {
	// implementasi query get user by id
	return []entities.User{}
}
