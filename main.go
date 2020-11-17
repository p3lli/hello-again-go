package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/p3lli/hello-again-go/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Error during env var loading: %s", err.Error)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! (...orld! ...rld! ... ld!)")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Port)))
}
