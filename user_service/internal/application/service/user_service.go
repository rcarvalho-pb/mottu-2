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
	"golang.org/x/crypto/bcrypt"
)

var allowedTypes = map[string]bool{
	"image/jpeg":      true,
	"image/jpg":       true,
	"application/pdf": true,
}

func (us *UserService) CreateUser(dto *dto.UserDTO) error {
	user := model.UserFromDTO(dto)
	if dto.CNHFileName != "" && dto.CNHFile != nil {
		contentType := http.DetectContentType(dto.CNHFile)
		if !allowedTypes[contentType] {
			return fmt.Errorf("invalid content type for cnh file. Only accept pdf")
		}
		cnhFilePath := getPathFromHash(generateHash(dto.CNHFileName))
		if err := os.MkdirAll(cnhFilePath, os.ModePerm); err != nil {
			return err
		}
		user.CNHFilePath = filepath.Join(cnhFilePath, fmt.Sprintf("%s_cnh.pdf", dto.Username))
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
		if err := os.MkdirAll(avatarFilePath, os.ModePerm); err != nil {
			return err
		}
		user.Avatar = filepath.Join(avatarFilePath, fmt.Sprintf("%s_avatar.jpeg", dto.Username))
		file, err := os.Create(user.Avatar)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := io.Copy(file, bytes.NewReader(dto.AvatarFile)); err != nil {
			return err
		}
	}
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := us.repository.CreateUser(user); err != nil {
		_ = os.RemoveAll(filepath.Join(baseDirectory, strings.Split(strings.TrimPrefix(user.Avatar, "/"), "/")[0]))
		_ = os.RemoveAll(filepath.Join(baseDirectory, strings.Split(strings.TrimPrefix(user.CNHFilePath, "/"), "/")[0]))
		return err
	}
	return nil
}

func (us *UserService) GetUserById(id int64) (*dto.UserDTO, error) {
	user, err := us.repository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	dto := user.ToDTO()
	return dto, err
}

func (us *UserService) GetUserByUsername(username string) (*dto.UserDTO, error) {
	user, err := us.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	dto := user.ToDTO()
	return dto, err
}

func (us *UserService) GetAllUsers() ([]*dto.UserDTO, error) {
	users, err := us.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	usersDTO := make([]*dto.UserDTO, 0)
	for _, u := range users {
		usersDTO = append(usersDTO, u.ToDTO())
	}
	return usersDTO, nil
}

func (us *UserService) GetAllActiveUsers() ([]*dto.UserDTO, error) {
	users, err := us.repository.GetAllUsers()
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
	user, err := us.repository.GetUserById(dto.Id)
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
	if dto.CNHFileName != "" && dto.CNHFile != nil {
		contentType := http.DetectContentType(dto.CNHFile)
		if !allowedTypes[contentType] {
			return fmt.Errorf("invalid content type for cnh file. Only accept pdf")
		}
		cnhFilePath := getPathFromHash(generateHash(dto.CNHFileName))
		oldFile := user.CNHFilePath
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
		if err := os.RemoveAll(filepath.Join(baseDirectory, strings.Split(strings.TrimPrefix(oldFile, "/"), "/")[0])); err != nil {
			return err
		}
	}
	if dto.AvatarFileName != "" && dto.AvatarFile != nil {
		contentType := http.DetectContentType(dto.CNHFile)
		if !allowedTypes[contentType] {
			return fmt.Errorf("invalid content type for cnh file. Only accept jpeg")
		}
		avatarFilePath := getPathFromHash(generateHash(dto.AvatarFileName))
		oldFile := user.Avatar
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
		if err := os.RemoveAll(filepath.Join(baseDirectory, strings.Split(strings.TrimPrefix(oldFile, "/"), "/")[0])); err != nil {
			return err
		}
	}
	user.UpdateTime()
	if err := us.repository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdatePassword(updatePassword *dto.UpdatePasswordDTO) error {
	user, err := us.repository.GetUserById(updatePassword.Id)
	if err != nil {
		return err
	}
	if err := comparePasswords(user.Password, updatePassword.Password); err != nil {
		return err
	}
	hashedPassword, err := hashPassword(updatePassword.NewPassword)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.UpdateTime()
	if err := us.repository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (us *UserService) DeactivateUser(id int64) error {
	user, err := us.repository.GetUserById(id)
	if err != nil {
		return err
	}
	user.Active = false
	user.UpdateTime()
	if err := us.repository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (us *UserService) ActivateUser(id int64) error {
	user, err := us.repository.GetUserById(id)
	if err != nil {
		return err
	}
	user.Active = true
	user.UpdateTime()
	if err := us.repository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (us *UserService) ChangeUserRole(userDTO *dto.UserDTO) error {
	user, err := us.repository.GetUserById(userDTO.Id)
	if err != nil {
		return err
	}
	u := model.UserFromDTO(userDTO)
	user.Role = u.Role
	if err := us.repository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (us *UserService) ValidatePassword(userDTO *dto.UserDTO) error {
	user, err := us.repository.GetUserById(userDTO.Id)
	if err != nil {
		return err
	}
	return comparePasswords(user.Password, userDTO.Password)
}

func (us *UserService) ComparePasswords(hashedPassword, password string) error {
	return comparePasswords(hashedPassword, password)
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func comparePasswords(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
