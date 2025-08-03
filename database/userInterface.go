package database

import (
	"goSre/entity"
)

type UserInterface interface {
	CreateUser(email string, senha string) (*entity.User, error)
	DeleteUser(email string) (*entity.User, error)
	UpdatePassword(email string, senha string) (*entity.User, error)
}
