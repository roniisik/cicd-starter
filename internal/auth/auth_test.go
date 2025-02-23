package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"basic test":        {input: "ApiKey 123", want: "123"},
		"empty header test": {input: "", want: ""},
		"no separator test": {input: "ApiKey123", want: ""},
		"wrong format test": {input: "Key 123", want: "1"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Set("Authorization", tc.input)
			got, _ := GetAPIKey(header)
			if tc.want != got {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
