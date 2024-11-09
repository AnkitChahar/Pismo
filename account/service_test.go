package account

import (
	"reflect"
	"testing"

	"pismo/models"
)

func TestService_CreateAccount(t *testing.T) {
	tests := []struct {
		name    string
		args    *models.Account
		want    *models.Account
		wantErr bool
	}{
		{
			name: "happy case",
			args: &models.Account{
				DocumentNumber: "12345",
			},
			want: &models.Account{
				ID:             1,
				DocumentNumber: "12345",
			},
			wantErr: false,
		},
		{
			name:    "error case",
			args:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: testDB,
			}
			got, err := s.CreateAccount(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAccountByID(t *testing.T) {
	mockAccount := &models.Account{
		ID:             5,
		DocumentNumber: "12345678",
	}

	insertMockData(mockAccount)

	tests := []struct {
		name    string
		args    uint
		want    *models.Account
		wantErr bool
	}{
		{
			name:    "happy case",
			args:    mockAccount.ID,
			want:    mockAccount,
			wantErr: false,
		},
		{
			name:    "record not found case",
			args:    3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				db: testDB,
			}
			got, err := s.GetAccountByID(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
