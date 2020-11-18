package config

import (
	"os"
	"strconv"
)

// RequestHandlerType custom type for request handler type
type RequestHandlerType string

// Available request handler types
const (
	HTTP RequestHandlerType = "net/http"
)

// ContentType custom type for content type
type ContentType string

// Content types
const (
	PNG             ContentType = "image/png"
	MOCKCONTENTTYPE             = "mock"
)

// DrawerType custom type for drawer type
type DrawerType string

// Available service type
const (
	GGDRAWER   DrawerType = "gg"
	MOCKDRAWER            = "mockDrawer"
)

// Config application container
type Config struct {
	Port               string
	ImgWidth           int
	ImgHeight          int
	RequestHandlerType RequestHandlerType
	ContentType        ContentType
	DrawerType         DrawerType
}

// LoadConfig loads configurations from environment variables
func LoadConfig() (*Config, error) {

	config := new(Config)

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

	config.RequestHandlerType = HTTP
	config.ContentType = PNG
	config.DrawerType = GGDRAWER

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
