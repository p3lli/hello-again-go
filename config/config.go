package config

import (
	"os"
	"strconv"
)

// Config application container
type Config struct {
	Port      string
	ImgWidth  int
	ImgHeight int
}

// LoadConfig loads configurations from environment variables
func LoadConfig() (*Config, error) {

	var config *Config

	var port string
	if port = os.Getenv("PORT"); port == "" {
		config.Port = "8080"
	} else {
		config.Port = port
	}

	imgWidthEnvIntValue, err := getEnvInt("IMG_WIDHT", 50)
	if err != nil {
		return nil, err
	}
	config.ImgWidth = imgWidthEnvIntValue

	imgHeightEnvIntValue, err := getEnvInt("IMG_HEIGHT", 50)
	if err != nil {
		return nil, err
	}
	config.ImgHeight = imgHeightEnvIntValue

	return config, nil
}

func getEnvInt(envName string, defaultValue int) (int, error) {
	envIntString := os.Getenv(envName)
	if envIntString == "" {
		return defaultValue, nil
	}
	envIntValue, err := strconv.Atoi(envIntString)
	if err != nil {
		return defaultValue, err
	}
	return envIntValue, nil
}
