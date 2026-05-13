package auth

import (
	"net/http"
	"testing"
	"os"
)

func TestGetAPIKey_Success(t *testing.T) {

	headers := http.Header{}
	headers.Add("Authorization", "ApiKey my-super-secret-key")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
	if key != "my-super-secret-key" {
		t.Errorf("expected 'my-super-secret-key', but got: %s", key)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {

	headers := http.Header{}

	key, err := GetAPIKey(headers)

	if err == nil {
		t.Errorf("expected an error because header is missing")
	}
	if key != "" {
		t.Errorf("expected empty string, but got: %s", key)
	}
}

func TestGetAPIKey_stringescape(t *testing.T) {

	headers := http.Header{}
	headers.Add("Authorization", "ApiKey '; DROP TABLE users;")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	if key != "';" {
		t.Errorf("expected %q, but got: %q", "';", key)
	}
//	os.Exit(1)
}
