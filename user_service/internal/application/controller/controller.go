package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/service"
)

type UserController struct {
	service.UserService
}

func New(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		ErrorJson(w, err)
		return
	}
	var userDTO dto.UserDTO
	jsonData := r.FormValue("json")
	if err := json.Unmarshal([]byte(jsonData), &userDTO); err != nil {
		ErrorJson(w, err, http.StatusUnprocessableEntity)
		return
	}
	avatar, avatarHeader, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		ErrorJson(w, fmt.Errorf("erro ao obter avatar [%s]", avatarHeader.Filename))
		return
	} else {
		avatarBytes, err := io.ReadAll(avatar)
		if err != nil {
			ErrorJson(w, fmt.Errorf("erro ao obter arquivo"))
			return
		}
		userDTO.AvatarFileName = avatarHeader.Filename
		// copy(userDTO.AvatarFile, avatarBytes)
		userDTO.AvatarFile = avatarBytes
	}
	defer avatar.Close()
	cnhFile, cnhHeader, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		ErrorJson(w, fmt.Errorf("erro ao obter avatar"))
		return
	} else {
		cnhBytes, err := io.ReadAll(cnhFile)
		if err != nil {
			ErrorJson(w, fmt.Errorf("erro ao obter arquivo"))
			return
		}
		userDTO.CNHFileName = cnhHeader.Filename
		// copy(userDTO.CNHFile, cnhBytes)
		userDTO.CNHFile = cnhBytes
	}
	defer cnhFile.Close()
	if err = uc.UserService.CreateUser(&userDTO); err != nil {
		ErrorJson(w, err)
		return
	}
	WriteJson(w, http.StatusCreated, nil)
}
