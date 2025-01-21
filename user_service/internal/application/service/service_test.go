package service

import (
	"testing"
	"time"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-user_service/internal/model"
	"github.com/stretchr/testify/mock"
)

// type UserRepository interface {
// 	CreateUser(*model.User) error
// 	UpdateUser(*model.User) error
// 	GetUserById(int64) (*model.User, error)
// 	GetUserByUsername(*string) (*model.User, error)
// 	GetAllUsers() ([]*model.User, error)
// }

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

func TestCreateUser(t *testing.T) error {
	userDTO := &dto.UserDTO{
		Username:       "Ramon",
		Password:       "123",
		Role:           "Admin",
		Name:           "Ramon",
		BirthDate:      time.Now(),
		CNPJ:           "123123123",
		CNH:            "124124124",
		CNHType:        "B",
		ActiveLocation: true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Active:         true,
		AvatarFileName: ,
		AvatarFile:     []byte{},
		CNHFileName:    "",
		CNHFile:        []byte{},
	}
	userRepository := new(UserRepositoryMock)
	userService := New(userRepository)
	userRepository.On("CreateUser")
}
