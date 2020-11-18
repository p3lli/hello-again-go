package draw

import (
	"fmt"
	"image"
)

// MockDrawer implements Drawer interface; used for tests
type MockDrawer struct {
	imgWidth  int
	imgHeight int
}

var _ Drawer = (*MockDrawer)(nil)

// DrawText return a gray coloured image; used for tests
func (m MockDrawer) DrawText(text string) (image.Image, error) {
	if text == "RETURN_ERROR" {
		return nil, fmt.Errorf("This is an error! Something went horribly wrong")
	}
	rect := image.Rect(0, 0, m.imgWidth, m.imgHeight)
	img := image.NewGray(rect)
	return img, nil
}
