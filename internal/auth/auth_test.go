package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {

	tests := []struct {
		name       string
		headers    http.Header
		wantApiKey string
		wantErr    bool
	}{
		{
			name:       "Empty Header",
			headers:    http.Header{},
			wantApiKey: "",
			wantErr:    true,
		},
		{
			name:       "Correct API Header",
			headers:    http.Header{"Authorization": []string{"ApiKeye gurbe"}},
			wantApiKey: "gurbe",
			wantErr:    false,
		},
		{
			name:       "Malformed Header",
			headers:    http.Header{"Authorization": []string{"ApiKaye gurbe gurbe"}},
			wantApiKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotApiKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApiKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotApiKey != tt.wantApiKey {
				t.Errorf("GetApiKey() error = %v, want %v", gotApiKey, tt.wantApiKey)
			}
		})
	}
}
