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
	h2 := http.Header{}
	h2.Set("Authorization", "")

	tests := []test{
		{
			input: h1,
			want: "somevalue",
		},
		{
			input: h2,
			want: "",
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if got != tc.want {
			t.Fatalf("expected: %v got: %v", tc.want, got)
		}
		if err != nil {
			t.Fatalf("expected no error: %v", err)
		}
	}
}
