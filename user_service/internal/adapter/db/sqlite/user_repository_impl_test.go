package sqlite_test

import (
	"testing"

	"github.com/rcarvalho-pb/mottu-user_service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	args := ur.Called(user)
	return args.Error(0)
}

func (ur *UserRepository) UpdateUser(user *model.User) error {
	args := ur.Called(user)
	return args.Error(0)
}

func (ur *UserRepository) GetUserById(id int64) (*model.User, error) {
	args := ur.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (ur *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	args := ur.Called(username)
	return args.Get(0).(*model.User), args.Error(1)
}

func (ur *UserRepository) GetAllUsers() ([]*model.User, error) {
	args := ur.Called()
	return args.Get(0).([]*model.User), args.Error(1)
}

func TestGetUserById(t *testing.T) {
	userRepository := new(UserRepository)
	user := &model.User{
		Id:       1,
		Username: "Teste",
		Password: "123",
	}
	userRepository.On("GetUserById", int64(1)).Return(user, nil)
	result, err := userRepository.GetUserById(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), result.Id)
	assert.Equal(t, "Teste", result.Username)
	userRepository.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	userRepository := new(UserRepository)
	user := &model.User{
		Username: "TesteCreate",
		Password: "senha123",
	}
	userRepository.On("CreateUser", user).Return(nil)
	err := userRepository.CreateUser(user)
	assert.NoError(t, err)
	userRepository.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	userRepository := new(UserRepository)
	user := &model.User{
		Id:       1,
		Username: "TesteUpdate",
		Password: "novaSenha123",
	}
	userRepository.On("UpdateUser", user).Return(nil)
	err := userRepository.UpdateUser(user)
	assert.NoError(t, err)
	userRepository.AssertExpectations(t)
}

func TestGetUserByUsername(t *testing.T) {
	userRepository := new(UserRepository)
	user := &model.User{
		Username: "TesteUsername",
		Password: "senha123",
	}
	userRepository.On("GetUserByUsername", "TesteUsername").Return(user, nil)
	result, err := userRepository.GetUserByUsername("TesteUsername")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "TesteUsername", result.Username)
	userRepository.AssertExpectations(t)
}

func TestGetAllUsers(t *testing.T) {
	userRepository := new(UserRepository)
	users := []*model.User{
		{
			Id:       1,
			Username: "User1",
			Password: "senha1",
		},
		{
			Id:       2,
			Username: "User2",
			Password: "senha2",
		},
	}
	userRepository.On("GetAllUsers").Return(users, nil)
	result, err := userRepository.GetAllUsers()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "User1", result[0].Username)
	assert.Equal(t, "User2", result[1].Username)
	userRepository.AssertExpectations(t)
}
