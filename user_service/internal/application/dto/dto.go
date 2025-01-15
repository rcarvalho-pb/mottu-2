package dto

import "time"

type UserDTO struct {
	Id             int64     `json:"id"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Role           string    `json:"role"`
	Name           string    `json:"name"`
	BirthDate      time.Time `json:"birth_date"`
	CNPJ           string    `json:"cnpj"`
	CNH            string    `json:"cnh"`
	CNHType        string    `json:"cnh_type"`
	CNHFilePath    []byte    `json:"cnh_file_path"`
	ActiveLocation bool      `json:"active_location"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Active         bool      `json:"active"`
}
