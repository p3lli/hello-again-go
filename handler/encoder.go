package handler

import (
	"bytes"
	"fmt"
	"hello-again-go/config"
	"image"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

// Encoder encodes an image and puts it in the response writer
type Encoder interface {
	WriteImage(w http.ResponseWriter, img *image.Image)
}

// NewEncoder is a factory method for Encoder implementations, based on writer "Content-Type"
func NewEncoder(conf config.Config) (Encoder, error) {
	if conf.ContentType == config.PNG {
		pngEncoder := new(PNGEncoder)
		pngEncoder.encoderContentType = conf.ContentType
		return pngEncoder, nil
	}
	return nil, fmt.Errorf("No encoder found for content type '%s'", conf.ContentType)
}

// PNGEncoder implements Encoder for png images
type PNGEncoder struct {
	encoderContentType config.ContentType
}

var _ Encoder = (*PNGEncoder)(nil)

// WriteImage encodes a png image in the response writer
func (PNGEncoder) WriteImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
