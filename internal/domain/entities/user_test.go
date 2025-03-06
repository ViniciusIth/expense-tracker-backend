package entities

import (
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
	}{
		{
			name:    "valid user",
			user:    User{Name: "John Doe", Email: "john@example.com"},
			wantErr: false,
		},
		{
			name:    "empty id",
			user:    User{ID: "", Name: "John Doe", Email: "john@example.com"},
			wantErr: true,
		},
		{
			name:    "empty name",
			user:    User{Name: "", Email: "john@example.com"},
			wantErr: true,
		},
		{
			name:    "empty email",
			user:    User{Name: "John Doe", Email: ""},
			wantErr: true,
		},
		{
			name:    "invalid email format",
			user:    User{Name: "John Doe", Email: "invalid-email"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
