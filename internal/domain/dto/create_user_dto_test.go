package dto

import (
	"testing"
)

func TestCreateUserRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     CreateUserRequest
		wantErr bool
	}{
		{
			name:    "valid request",
			req:     CreateUserRequest{Name: "John Doe", Email: "john@example.com", Password: "password123"},
			wantErr: false,
		},
		{
			name:    "empty name",
			req:     CreateUserRequest{Name: "", Email: "john@example.com", Password: "password123"},
			wantErr: true,
		},
		{
			name:    "empty email",
			req:     CreateUserRequest{Name: "John Doe", Email: "", Password: "password123"},
			wantErr: true,
		},
		{
			name:    "empty password",
			req:     CreateUserRequest{Name: "John Doe", Email: "john@example.com", Password: ""},
			wantErr: true,
		},
		{
			name:    "password too short",
			req:     CreateUserRequest{Name: "John Doe", Email: "john@example.com", Password: "short"},
			wantErr: true,
		},
	}

	for _, example := range tests {
		t.Run(example.name, func(t *testing.T) {
			err := example.req.Validate()
			if (err != nil) != example.wantErr {
				t.Errorf("CreateUserRequest.Validate() error = %v, wantErr %v", err, example.wantErr)
			}
		})
	}
}
