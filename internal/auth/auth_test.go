package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_EmptyHeader(t *testing.T) {
	key, err := GetAPIKey(http.Header{})
	if key != "" || err == nil {
		t.Errorf("empty input should give empty API key")
	}
}

func TestGetAPIKey_NoAuthorizationHeader(t *testing.T) {
	h := http.Header{}
	h.Add("Some different header who cares", "value")
	key, err := GetAPIKey(h)
	if key != "" || err == nil {
		t.Errorf("empty input should give empty API key")
	}
}

func TestGetAPIKey_WrongAuthorization(t *testing.T) {
	h := http.Header{}
	h.Add("Authorization", "wrong auth")
	key, err := GetAPIKey(h)
	if key != "" || err == nil {
		t.Errorf("empty input should give empty API key")
	}
}

func TestGetAPIKey_GoodHeader(t *testing.T) {
	h := http.Header{}
	h.Add("Authorization", "ApiKey 543")
	key, err := GetAPIKey(h)
	if key != "543" || err != nil {
		t.Errorf("empty input should give empty API key")
	}
}
