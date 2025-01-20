package service

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/repository"
	"github.com/rcarvalho-pb/mottu-user_service/internal/model"
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

func (us *UserService) CreateUser(dto *dto.UserDTO) error {
	user := model.UserFromDTO(dto)
	if dto.CNHFileName != "" && dto.CNHFile != nil {
		cnhFilePath := getPathFromHash(generateHash(dto.CNHFileName))
		user.CNHFilePath = filepath.Join(cnhFilePath, dto.CNHFileName)
		if err := os.MkdirAll(cnhFilePath, os.ModePerm); err != nil {
			return err
		}
		file, err := os.Create(user.CNHFilePath)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err = io.Copy(file, dto.CNHFile); err != nil {
			return err
		}
	}
	if dto.AvatarFileName != "" && dto.AvatarFile != nil {
		avatarFilePath := filepath.Join(getPathFromHash(generateHash(dto.AvatarFileName)), dto.AvatarFileName)
		user.Avatar = avatarFilePath
	}

	if err := us.UserRepository.CreateUser(model.UserFromDTO(dto)); err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserById(id int64) (*dto.UserDTO, error) {
	user, err := us.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	dto := user.ToDTO()
	return dto, err
}

func (us *UserService) UpdateUser(dto *dto.UserDTO) error {
	user, err := us.GetUserById(dto.Id)
	if err != nil {
		return err
	}
	if dto.Username != "" {
		user.Username = dto.Username
	}
	if dto.Name != "" {
		user.Name = dto.Name
	}
	if !dto.BirthDate.IsZero() {
		user.BirthDate = dto.BirthDate
	}
	if dto.CNH != "" {
		user.CNH = dto.CNH
	}
	if dto.CNPJ != "" {
		user.CNPJ = dto.CNPJ
	}
	if dto.CNHType != "" {
		user.CNHType = dto.CNHType
	}
	return nil
}
