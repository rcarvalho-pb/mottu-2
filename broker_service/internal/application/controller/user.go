package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/helper"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/service"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

type UserController struct {
	service *service.Service
}

func newUserController(serv *service.Service) UserController {
	return UserController{service: serv}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Broker: received new user creation request")
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		helper.ErrorJson(w, err)
		return
	}
	var userDTO model.UserDTO
	jsonData := r.FormValue("json")
	if err := json.Unmarshal([]byte(jsonData), &userDTO); err != nil {
		helper.ErrorJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	avatar, avatarHeader, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		helper.ErrorJson(w, fmt.Errorf("erro ao obter avatar [%s]", avatarHeader.Filename))
		return
	} else {
		avatarBytes, err := io.ReadAll(avatar)
		if err != nil {
			helper.ErrorJson(w, fmt.Errorf("erro ao obter arquivo"))
			return
		}
		userDTO.AvatarFileName = avatarHeader.Filename
		copy(userDTO.AvatarFile, avatarBytes)
		userDTO.AvatarFile = avatarBytes
	}
	defer avatar.Close()
	cnhFile, cnhHeader, err := r.FormFile("cnh")
	if err != nil && err != http.ErrMissingFile {
		helper.ErrorJson(w, fmt.Errorf("erro ao obter avatar"))
		return
	} else {
		cnhBytes, err := io.ReadAll(cnhFile)
		if err != nil {
			helper.ErrorJson(w, fmt.Errorf("erro ao obter arquivo"))
			return
		}
		userDTO.CNHFileName = cnhHeader.Filename
		copy(userDTO.CNHFile, cnhBytes)
		userDTO.CNHFile = cnhBytes
	}
	defer cnhFile.Close()
	if err = uc.service.UserService.CreateUser(&userDTO); err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusCreated, nil)
	log.Println("Broker: user successfully saved")
}

func (uc *UserController) GetAllActiveUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Broker: get all active users request")
	users, err := uc.service.UserService.GetAllActiveUsers()
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, users)
	log.Println("Broker: all active users returned")
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Broker: get all users request")
	users, err := uc.service.UserService.GetAllUsers()
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, users)
	log.Println("Broker: all users returned")
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {

}
