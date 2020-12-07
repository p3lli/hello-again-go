package handler

import (
	"fmt"
	"hello-again-go/config"
	"hello-again-go/service/draw"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
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
	text, ok := query["text"]
	if !ok {
		log.WithFields(
			log.Fields{
				"error": "Query param 'text' not provided",
			}).Error("Query param 'text' not provided")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Query param 'text' not provided"))
		return
	}
	img, err := r.drawer.DrawText(strings.Join(text, ""))
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err.Error(),
			}).Errorf("Error during text '%s' drawing: %s", text, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("400 - Error during text '%s' drawing", text)))
		return
	}
	r.encoder.WriteImage(w, &img)
}
