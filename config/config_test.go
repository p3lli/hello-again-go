package config

import "testing"

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		t.Errorf(
			"It should not return errors; got '%s'",
			err.Error(),
		)
	}
	if config == nil {
		t.Error("Config container should not be nil")
	}
	if config != nil && config.Port != "8080" {
		t.Errorf(
			"Port value should be default value 8080; got '%s'",
			config.Port,
		)
	}
	if config != nil && config.ImgWidth != 50 {
		t.Errorf(
			"Image width value should be default value 50; got '%d'",
			config.ImgWidth,
		)
	}
	if config != nil && config.ImgHeight != 50 {
		t.Errorf(
			"Image height value should be default value 50; got '%d'",
			config.ImgHeight,
		)
	}
}
