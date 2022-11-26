package repo

import (
	"encoding/json"
	"io/fs"
	"os"
	"strconv"

	"github.com/torikki-tou/trueconf_testtask/config"
	"github.com/torikki-tou/trueconf_testtask/entity"
	"github.com/torikki-tou/trueconf_testtask/common/custom_error"
)

type UserRepositiry interface {
	GetUsers() (entity.UserList, error)
	GetUserByID(userID string) (entity.User, error) 
	InsertUser(new_obj entity.User) (int, error)
	UpdateUser(userID string, new_obj entity.User) error
	DeleteUser(userID string) error
}

func NewUserRepository(filename string) UserRepositiry {
	return &userRepositiry{
		filename: filename,
	}
}

type userRepositiry struct {
	filename string
}

func (r *userRepositiry) getUserStore() (*entity.UserStore, error) {
	f, err := os.ReadFile(config.Filename)
	if err != nil {
		return nil, err
	}

	s := entity.UserStore{}
	_ = json.Unmarshal(f, &s)
	return &s, nil
}

func (r *userRepositiry) commit(store *entity.UserStore) {
	b, _ := json.MarshalIndent(store, "", " ")
	_ = os.WriteFile(config.Filename, b, fs.ModePerm)
}

func (r *userRepositiry) GetUsers() (entity.UserList, error) {
	store, err := r.getUserStore()
	if err != nil {
		return nil, err
	}

	return store.List, nil
}

func (r *userRepositiry) GetUserByID(userID string) (entity.User, error) {
	store, err := r.getUserStore()
	if err != nil {
		return entity.User{}, err
	}

	user, ok := store.List[userID]
	if !ok {
		return entity.User{}, &custom_error.ErrObjectNotFound{}
	}

	return user, nil
}

func (r *userRepositiry) InsertUser(new_obj entity.User) (int, error) {
	store, err := r.getUserStore()
	if err != nil {
		return 0, err
	}

	store.Increment++

	id := strconv.Itoa(store.Increment)
	store.List[id] = new_obj

	r.commit(store)

	return store.Increment, nil
}

func (r *userRepositiry) UpdateUser(userID string, new_obj entity.User) error {
	store, err := r.getUserStore()
	if err != nil {
		return err
	}

	if _, ok := store.List[userID]; !ok {
		return &custom_error.ErrObjectNotFound{}
	}

	store.List[userID] = new_obj

	r.commit(store)
	return nil
}

func (r *userRepositiry) DeleteUser(userID string) error {
	store, err := r.getUserStore()
	if err != nil {
		return err
	}

	if _, ok := store.List[userID]; !ok {
		return &custom_error.ErrObjectNotFound{}
	}

	delete(store.List, userID)

	r.commit(store)
	return nil
}