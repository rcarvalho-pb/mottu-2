package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/global"
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
}

func (uc *UserController) GetAllActiveUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.service.UserService.GetAllActiveUsers()
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, users)
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.service.UserService.GetAllUsers()
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, users)
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		helper.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	user, err := uc.service.UserService.GetUserByiD(userId)
	if err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(global.CLAIMS).(model.Claims)
	if !ok {
		helper.ErrorJson(w, fmt.Errorf("erro recuperando context"))
		return
	}
	userId, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		helper.ErrorJson(w, err, http.StatusBadRequest)
		return
	}
	if userId != claims.UserId && claims.UserRole != "admin" {
		helper.ErrorJson(w, fmt.Errorf("only admin can update other users"), http.StatusUnauthorized)
		return
	}
	var userDTO model.UserDTO
	if err := helper.ReadJson(w, r, &userDTO); err != nil {
		helper.ErrorJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	userDTO.Id = userId
	if err := uc.service.UserService.UpdateUser(&userDTO); err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, nil)
}

func (uc *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(global.CLAIMS).(model.Claims)
	if !ok {
		helper.ErrorJson(w, fmt.Errorf("erro recuperando context"))
		return
	}
	userId, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	if userId != claims.UserId && claims.UserRole != "admin" {
		helper.ErrorJson(w, fmt.Errorf("only admin can update other users"), http.StatusUnauthorized)
		return
	}
	var newPassword model.UpdatePasswordDTO
	if err := helper.ReadJson(w, r, &newPassword); err != nil {
		helper.ErrorJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	newPassword.Id = userId
	if err := uc.service.UserService.UpdatePassword(&newPassword); err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, nil)
}

func (uc *UserController) DeactivateUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(global.CLAIMS).(model.Claims)
	if !ok {
		helper.ErrorJson(w, fmt.Errorf("erro recuperando context"))
		return
	}
	userId, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	if userId != claims.UserId && claims.UserRole != "admin" {
		helper.ErrorJson(w, fmt.Errorf("only admin can update other users"), http.StatusUnauthorized)
		return
	}
	if err := uc.service.UserService.DeactivateUser(userId); err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, nil)
}

func (uc *UserController) ReactivateUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(global.CLAIMS).(model.Claims)
	if !ok {
		helper.ErrorJson(w, fmt.Errorf("erro recuperando context"))
		return
	}
	userId, err := strconv.ParseInt(r.PathValue("userId"), 10, 64)
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	if userId != claims.UserId && claims.UserRole != "admin" {
		helper.ErrorJson(w, fmt.Errorf("only admin can update other users"), http.StatusUnauthorized)
		return
	}
	if err := uc.service.UserService.ReactivateUser(userId); err != nil {
		helper.ErrorJson(w, err, http.StatusInternalServerError)
		return
	}
	helper.WriteJson(w, http.StatusOK, nil)
}
