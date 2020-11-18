package draw

import (
	"fmt"
	"hello-again-go/config"
	"image"

	"github.com/fogleman/gg"
)

// Drawer interface
type Drawer interface {
	DrawText(text string) (image.Image, error)
}

// GGDrawer implements Drawer interface
type GGDrawer struct {
	imgWidth  int
	imgHeight int
}

var _ Drawer = (*GGDrawer)(nil)

// NewDrawer return an implementation of Drawer based on 'drawerType' value
func NewDrawer(conf config.Config) (Drawer, error) {
	if conf.DrawerType == config.GGDRAWER {
		ggDrawer := new(GGDrawer)
		ggDrawer.imgWidth = conf.ImgHeight
		ggDrawer.imgHeight = conf.ImgWidth
		return ggDrawer, nil
	}
	if conf.DrawerType == config.MOCKDRAWER {
		mockDrawer := new(MockDrawer)
		mockDrawer.imgWidth = conf.ImgHeight
		mockDrawer.imgHeight = conf.ImgWidth
		return mockDrawer, nil
	}
	return nil, fmt.Errorf("No drawer found for type '%s'", conf.DrawerType)
}

// DrawText return a gray coloured image
// with the passed text written at the center
func (d GGDrawer) DrawText(text string) (image.Image, error) {
	dc := gg.NewContext(d.imgWidth, d.imgHeight)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(text,
		float64(d.imgWidth/2),
		float64(d.imgHeight/2),
		0.5,
		0.5,
	)
	return dc.Image(), nil
}
