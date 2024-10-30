package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	h1 := http.Header{}
	h1.Set("Authorization", "ApiKey somevalue")

	tests := []test{
		{
			input: h1,
			want:  "somevalue",
		},
	}

	for _, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if got != tc.want {
			t.Fatalf("expected: %v got: %v", tc.want, got)
		}
	}
}
