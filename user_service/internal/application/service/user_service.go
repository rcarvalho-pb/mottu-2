package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-user_service/internal/model"
)

func (us *UserService) CreateUser(dto *dto.UserDTO) error {
	allowedTypes := map[string]bool{
		"image/jpeg":          true,
		"application/pdf": true,
	}
	user := model.UserFromDTO(dto)
	if dto.CNHFileName != "" && dto.CNHFile != nil {
		contentType := http.DetectContentType(dto.CNHFile)
		if !allowedTypes[contentType] {
			return fmt.Errorf("invalid content type for cnh file. Only accept pdf")
		}
		cnhFilePath := getPathFromHash(generateHash(dto.CNHFileName))
		user.CNHFilePath = filepath.Join(cnhFilePath, fmt.Sprintf("%s_cnh.pdf", dto.Username))
		if err := os.MkdirAll(cnhFilePath, os.ModePerm); err != nil {
			return err
		}
		file, err := os.Create(user.CNHFilePath)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err = io.Copy(file, bytes.NewReader(dto.CNHFile)); err != nil {
			return err
		}
	}
	if dto.AvatarFileName != "" && dto.AvatarFile != nil {
		contentType := http.DetectContentType(dto.CNHFile)
		if !allowedTypes[contentType] {
			return fmt.Errorf("invalid content type for cnh file. Only accept jpeg")
		}
		avatarFilePath := getPathFromHash(generateHash(dto.AvatarFileName))
		user.Avatar = filepath.Join(avatarFilePath, fmt.Sprintf("%s_avatar.jpeg", dto.Username))
		if err := os.MkdirAll(avatarFilePath, os.ModePerm); err != nil {
			return err
		}
		file, err := os.Create(user.Avatar)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := io.Copy(file, bytes.NewReader(dto.AvatarFile)); err != nil {
			return err
		}
	}
	if err := us.UserRepository.CreateUser(model.UserFromDTO(dto)); err != nil {
		_ = os.RemoveAll(strings.Split(user.Avatar, "/")[0])
		_ = os.RemoveAll(strings.Split(user.CNHFilePath, "/")[0])
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

func (us *UserService) GetAllUsers() ([]*dto.UserDTO, error) {
	users, err := us.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	usersDTO := make([]*dto.UserDTO, 0)
	for _, u := range users {
		if u.Active {
			usersDTO = append(usersDTO, u.ToDTO())
		}
	}
	return usersDTO, nil
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
