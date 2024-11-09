package transaction

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"pismo/models"
)

func TestService_CreateTransaction(t *testing.T) {

	tests := []struct {
		name     string
		args     *models.Transaction
		want     *models.Transaction
		wantErr  bool
		mockFunc func(m *MockDeps)
	}{
		{
			name: "happy case",
			args: &models.Transaction{
				AccountID:       1,
				OperationTypeID: 4,
				Amount:          123,
			},
			want: &models.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: 4,
				Amount:          123,
				EventDate:       time.Now(),
			},
			wantErr: false,
			mockFunc: func(m *MockDeps) {
				m.AccountSvc.EXPECT().GetAccountByID(uint(1)).Return(nil, nil)
			},
		},
		{
			name: "account not found case",
			args: &models.Transaction{
				AccountID:       1,
				OperationTypeID: 4,
				Amount:          123,
			},
			wantErr: true,
			mockFunc: func(m *MockDeps) {
				m.AccountSvc.EXPECT().GetAccountByID(uint(1)).Return(nil, fmt.Errorf("account not found"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, m := setupServiceWithMocks(t)

			if tt.mockFunc != nil {
				tt.mockFunc(m)
			}

			got, err := s.CreateTransaction(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err == nil {
					t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}

			// Ignore the EventDate field as it is set by time.Now()
			got.EventDate = tt.want.EventDate
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
