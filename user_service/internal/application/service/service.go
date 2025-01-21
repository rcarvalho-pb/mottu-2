package service

import (
	"os"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/repository"
)

type UserService struct {
	repository.UserRepository
}

func New(rep repository.UserRepository) *UserService {
	baseDirectory = os.Getenv("IMAGES_DIRECTORY")
	return &UserService{
		UserRepository: rep,
	}
}
