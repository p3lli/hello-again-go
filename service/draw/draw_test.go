package draw

import (
	"hello-again-go/config"
	"testing"
)

// Content type used only for tests
const NOTVALID config.DrawerType = "notValid"

func TestNewDrawer(t *testing.T) {
	cases := map[string]struct {
		drawerType config.DrawerType
		hasError   bool
	}{
		"service": {
			drawerType: config.GGDRAWER,
			hasError:   false,
		},
		"mock": {
			drawerType: config.MOCKDRAWER,
			hasError:   false,
		},
		"not_valid": {
			drawerType: NOTVALID,
			hasError:   true,
		},
	}

	config, _ := config.LoadConfig()

	for k, c := range cases {
		config.DrawerType = c.drawerType
		encoder, err := NewDrawer(*config)
		if encoder != nil && c.hasError {
			t.Errorf("Test case: '%s': expected error, got none", k)
		}
		if err != nil && !c.hasError {
			t.Errorf("Test case: '%s': expected no errors, got '%s'", k, err.Error())
		}
	}
}
