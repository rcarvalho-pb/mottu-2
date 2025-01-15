package model_test

import (
	"testing"
	"time"

	. "github.com/rcarvalho-pb/mottu-user_service/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetUserRoleString(t *testing.T) {
	birthDate, _ := time.Parse("01/14/2023", "22/07/2024")
	createdAt, _ := time.Parse("01/14/2023", "01/01/2025")
	user := &User{
		Id:             1,
		Username:       "Papelito",
		Password:       "123",
		Role:           0,
		Name:           "Ramon",
		BirthDate:      birthDate,
		CNPJ:           "",
		CNH:            "",
		CNHType:        "",
		CNHFilePath:    "",
		ActiveLocation: false,
		CreatedAt:      createdAt,
		UpdatedAt:      createdAt,
		Active:         false,
	}

	userDTO := user.ToDTO()

	assert.Equal(t, user.Username, userDTO.Username)
	assert.Equal(t, userDTO.Role, "admin")
}
