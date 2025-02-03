package model

import "time"

type UserDTO struct {
	Id             int64     `json:"id,omitempty"`
	Username       string    `json:"username,omitempty"`
	Password       string    `json:"password,omitempty"`
	Role           string    `json:"role,omitempty"`
	Name           string    `json:"name,omitempty"`
	BirthDate      time.Time `json:"birth_date,omitempty"`
	CNPJ           string    `json:"cnpj,omitempty"`
	CNH            string    `json:"cnh,omitempty"`
	CNHType        string    `json:"cnh_type,omitempty"`
	ActiveLocation bool      `json:"active_location,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	Active         bool      `json:"active,omitempty"`
	AvatarFileName string
	AvatarFile     []byte
	CNHFileName    string
	CNHFile        []byte
}
