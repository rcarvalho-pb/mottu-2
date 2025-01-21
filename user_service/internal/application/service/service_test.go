package service_test

import (
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-user_service/internal/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) CreateUser(user *model.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserRepositoryMock) UpdateUser(user *model.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserRepositoryMock) GetUserById(id int64) (*model.User, error) {
	args := u.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (u *UserRepositoryMock) GetUserByUsername(username string) (*model.User, error) {
	args := u.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (u *UserRepositoryMock) GetAllUsers() ([]*model.User, error) {
	args := u.Called()
	return args.Get(0).([]*model.User), args.Error(1)
}

func TestCreateUser(userDTO *dto.UserDTO) error {
	userRepository := new(UserRepositoryMock)
	userService := service.New(userRepository)
}
