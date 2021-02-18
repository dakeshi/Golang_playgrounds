package main

import "testing"

func TestFetchJson(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FetchJson(); got != tt.want {
				t.Errorf("FetchJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
