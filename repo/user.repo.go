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

type UserRepository interface {
	GetUsers() (entity.UserList, error)
	GetUserByID(userID string) (entity.User, error) 
	InsertUser(new_obj entity.User) (int, error)
	UpdateUser(userID string, new_obj entity.User) error
	DeleteUser(userID string) error
}

func NewUserRepository(filename string) UserRepository {
	return &userRepository{
		filename: filename,
	}
}

type userRepository struct {
	filename string
}

func (r *userRepository) getUserStore() (*entity.UserStore, error) {
	f, err := os.ReadFile(config.Filename)
	if err != nil {
		return nil, err
	}

	s := entity.UserStore{}
	_ = json.Unmarshal(f, &s)
	return &s, nil
}

func (r *userRepository) commit(store *entity.UserStore) {
	b, _ := json.MarshalIndent(store, "", " ")
	_ = os.WriteFile(config.Filename, b, fs.ModePerm)
}

func (r *userRepository) GetUsers() (entity.UserList, error) {
	store, err := r.getUserStore()
	if err != nil {
		return nil, err
	}

	return store.List, nil
}

func (r *userRepository) GetUserByID(userID string) (entity.User, error) {
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

func (r *userRepository) InsertUser(new_obj entity.User) (int, error) {
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

func (r *userRepository) UpdateUser(userID string, new_obj entity.User) error {
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

func (r *userRepository) DeleteUser(userID string) error {
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