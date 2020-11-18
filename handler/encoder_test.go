package handler

import (
	"hello-again-go/config"
	"testing"
)

// Content type used only for tests
const NOTVALID config.ContentType = "not/valid"

func TestNewEncoder(t *testing.T) {
	cases := map[string]struct {
		contentType config.ContentType
		hasError    bool
	}{
		"png": {
			contentType: config.PNG,
			hasError:    false,
		},
		"not_valid": {
			contentType: NOTVALID,
			hasError:    true,
		},
	}

	config, _ := config.LoadConfig()

	for k, c := range cases {
		config.ContentType = c.contentType
		encoder, err := NewEncoder(*config)
		if encoder != nil && c.hasError {
			t.Errorf("Test case: '%s': expected error, got none", k)
		}
		if err != nil && !c.hasError {
			t.Errorf("Test case: '%s': expected no errors, got '%s'", k, err.Error())
		}
	}
}
