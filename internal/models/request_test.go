package models

import "testing"

func TestOrderRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request OrderRequest
		wantErr bool
	}{
		{
			name:    "Valid Request",
			request: OrderRequest{UserID: 1, ProductID: 100},
			wantErr: false,
		},
		{
			name:    "Invalid User ID (Zero)",
			request: OrderRequest{UserID: 0, ProductID: 100},
			wantErr: true,
		},
		{
			name:    "Invalid User ID (Negative)",
			request: OrderRequest{UserID: -5, ProductID: 100},
			wantErr: true,
		},
		{
			name:    "Invalid Product ID",
			request: OrderRequest{UserID: 1, ProductID: 0},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()

			// Do we expect errors? (tt.wantErr)
			// Do an error really occurs? (err != nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
