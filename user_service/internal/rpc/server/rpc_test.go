package rpc_server_test

import (
	"testing"

	"github.com/rcarvalho-pb/mottu-user_service/internal/application/dto"
	"github.com/rcarvalho-pb/mottu-user_service/internal/application/service"
)

func TestRPCServer_ValidatePassword(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		service *service.UserService
		port    string
		// Named input parameters for target function.
		userDTO *dto.UserDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rpc_server.New(tt.service, tt.port)
			gotErr := r.ValidatePassword(tt.userDTO, nil)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ValidatePassword() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ValidatePassword() succeeded unexpectedly")
			}
		})
	}
}
