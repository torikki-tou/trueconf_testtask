package repo

import "github.com/torikki-tou/trueconf_testtask/entity"

type UserRepositiry interface {
	GetUser() (entity.User, error) 
	InsertUser() (entity.User, error)
	UpdateUser() error
	DeleteUser() error
}

func NewUserRepository(filename string) UserRepositiry {
	return &userRepositiry{
		filename: filename,
	}
}

type userRepositiry struct {
	filename string
}

func (r *userRepositiry) GetUser() (entity.User, error) {return entity.User{}, nil}
func (r *userRepositiry) InsertUser() (entity.User, error) {return entity.User{}, nil}
func (r *userRepositiry) UpdateUser() error {return nil}
func (r *userRepositiry) DeleteUser() error {return nil}