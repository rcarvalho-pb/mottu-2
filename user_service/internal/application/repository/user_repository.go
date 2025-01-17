package repository

import "github.com/rcarvalho-pb/mottu-user_service/internal/model"

type UserRepository interface {
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	GetUserById(int64) (*model.User, error)
	GetUserByUsername(*string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
}
