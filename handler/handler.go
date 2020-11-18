package handler

import (
	"fmt"
	"hello-again-go/config"
	"hello-again-go/service/draw"
	"net/http"
	"strings"
)

// RequestHandler defines an interface for request handler
type RequestHandler interface {
	RespondImage(w http.ResponseWriter, r *http.Request)
}

// HTTPRequestHandler implements RequestHandler
type HTTPRequestHandler struct {
	encoder Encoder
	drawer  draw.Drawer
}

var _ RequestHandler = (*HTTPRequestHandler)(nil)

// NewRequestHandler returns an implementation of RequestHandler based on request handler type
func NewRequestHandler(conf config.Config) (RequestHandler, error) {
	if conf.RequestHandlerType == config.HTTP {
		httpRequestHandler := new(HTTPRequestHandler)

		encoder, err := NewEncoder(conf)
		if err != nil {
			return nil, err
		}
		httpRequestHandler.encoder = encoder

		drawer, err := draw.NewDrawer(conf)
		if err != nil {
			return nil, err
		}
		httpRequestHandler.drawer = drawer

		return httpRequestHandler, nil
	}
	return nil, fmt.Errorf("No request handler found for type '%s'", conf.RequestHandlerType)
}

// RespondImage return an image as response
func (r HTTPRequestHandler) RespondImage(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	text, _ := query["text"]
	img, err := r.drawer.DrawText(strings.Join(text, ""))
	if err != nil {
		panic(err)
	}
	r.encoder.WriteImage(w, &img)
}
