package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	testApiKey := "ApiKey 12345"
	header := http.Header{}
	header.Add("Authorization", testApiKey)
	actualAPIKey, err := auth.GetAPIKey(header)

	if err != nil {
		t.Fatalf("apikey err: %v", err)
	}

	expectedAPIKey := "12345"
	if actualAPIKey != expectedAPIKey {
		t.Fatalf("API key - expected: %s, actual %s", expectedAPIKey, actualAPIKey)
	}
}
