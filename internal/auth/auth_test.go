package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "ApiKey 1234")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("errored: %v", err)
	}
	if apiKey != `1234` {
		t.Fatalf("expected: %v, got: %v", `1234`, apiKey)
	}
}
