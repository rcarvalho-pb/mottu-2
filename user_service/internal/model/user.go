package model

import (
	"slices"
	"time"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
)

type Role int64

var ROLES = []string{"admin", "common"}

const (
	ADMIN Role = iota
	COMMON
)

func (r Role) String() string {
	return ROLES[r]
}

func getRoleCod(role string) Role {
	return Role(slices.Index(ROLES, role))
}

type UserRepository interface {
	CreateUser(*User) error
	UpdateUser(*User) error
	GetUserById(*int64) (*User, error)
	GetUserByUsername(*string) (*User, error)
	GetAllUsers() ([]*User, error)
}

type User struct {
	Id             int64     `db:"id"`
	Username       string    `db:"username"`
	Password       string    `db:"password"`
	Role           Role      `db:"role"`
	Name           string    `db:"name"`
	BirthDate      time.Time `db:"birth_date"`
	CNPJ           string    `db:"cnpj"`
	CNH            string    `db:"cnh"`
	CNHType        string    `db:"cnh_type"`
	CNHFilePath    string    `db:"cnh_file_path"`
	ActiveLocation bool      `db:"active_location"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	Active         bool      `db:"active"`
}

func (u *User) UpdateTime() {
	u.UpdatedAt = time.Now()
}

func (u *User) ToDTO() *dto.UserDTO {
	return &dto.UserDTO{
		Id:             u.Id,
		Username:       u.Username,
		Password:       u.Password,
		Role:           u.Role.String(),
		Name:           u.Name,
		BirthDate:      u.BirthDate,
		CNPJ:           u.CNPJ,
		CNH:            u.CNH,
		CNHType:        u.CNHType,
		CNHFilePath:    []byte{},
		ActiveLocation: u.ActiveLocation,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
		Active:         u.Active,
	}
}

func UserFromDTO(dto *dto.UserDTO) *User {
	return &User{
		Id:             dto.Id,
		Username:       dto.Username,
		Password:       dto.Password,
		Role:           getRoleCod(dto.Role),
		Name:           dto.Name,
		BirthDate:      dto.BirthDate,
		CNPJ:           dto.CNPJ,
		CNH:            dto.CNH,
		CNHType:        dto.CNHType,
		CNHFilePath:    "",
		ActiveLocation: dto.ActiveLocation,
		CreatedAt:      dto.CreatedAt,
		UpdatedAt:      dto.UpdatedAt,
		Active:         dto.Active,
	}
}
