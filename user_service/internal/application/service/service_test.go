package service

import (
	"fmt"
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

func TestCreateUser(t *testing.T) {
	userDTO := &dto.UserDTO{
		Username:       "Ramon",
		Password:       "123",
		Role:           "admin",
		Name:           "Ramon",
		BirthDate:      time.Now(),
		CNPJ:           "123123",
		CNH:            "12341234",
		CNHType:        "B",
		ActiveLocation: false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Active:         true,
		AvatarFileName: "",
		AvatarFile:     nil,
		CNHFileName:    "",
		CNHFile:        nil,
	}
	userRepository := new(UserRepositoryMock)
	fmt.Println("teste")
	userService := New(userRepository)
	userRepository.On("CreateUser", mock.Anything).Return(nil)
	if err := userService.CreateUser(userDTO); err != nil {
		t.Errorf(err.Error())
	}
}
