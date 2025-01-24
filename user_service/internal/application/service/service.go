package service

import (
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func New(rep repository.UserRepository, imagesDirectory string) *UserService {
	baseDirectory = imagesDirectory
	return &UserService{
		repository: rep,
	}
}
