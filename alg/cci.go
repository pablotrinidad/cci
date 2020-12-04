// package alg contains a cloud cover index calculator that performs the computations using a source and a mask image.
package alg

import (
	"fmt"
	"image"
	"image/color"
)

// CCI is a Cloud Coverage Index calculator.
type CCI struct {
	src, mask image.Image
	seg       image.Image
	out       *image.RGBA
}

// NewCCI returns a new Cloud Coverage Index calculator instance. It expects the source to be a sky photography while
// the mask to be a black RGBA: (0,0,0,2255) and white RGBA: (255,255,255,255) image. It will align the mask and the
// source image in the center and only process those pixels that lay in the intersection of both images and that map to
// a white pixel in the mask. All black pixels will be discarded out of the computation.
func NewCCI(src, mask image.Image) *CCI {
	cci := &CCI{src: src, mask: mask}
	return cci
}

// Run returns the cloud cover index obtained from the source image applying the mask.
// Cloud Cover Index is calculated tagging individual pixels as cloud or sky using their RGB value.
// Pixels are considered cloud when they R/B ratio is >= 0.95, otherwise
func (c *CCI) Run() float64 {
	outputBounds := c.OutputBounds()
	maskOffset, srcOffset := c.getImagesOffset()

	c.out = image.NewRGBA(outputBounds)
	area := 0.0
	clouds := 0.0

	for y := 0; y < outputBounds.Max.Y; y++ {
		for x := 0; x < outputBounds.Max.X; x++ {
			maskColor := c.mask.At(x+maskOffset.X, y+maskOffset.Y)
			srcColor := c.src.At(x+srcOffset.X, y+srcOffset.Y)
			if isWhite(maskColor) {
				r, _, b, _ := srcColor.RGBA()
				if float64(r)/float64(b) < 0.95 {
					c.out.Set(x, y, color.Black)
				} else {
					c.out.Set(x, y, color.White)
					clouds++
				}
				area++
			}
		}
	}
	return clouds / area
}

// SaveSegmentation returns the image that resulted from the cloud segmentation algorithm.
// Run() has to be called before this method such that output result is ready.
func (c *CCI) SaveSegmentation() (*image.RGBA, error) {
	if c.out == nil {
		return nil, fmt.Errorf("CCI have not been calculated yet, please call Run() first")
	}
	return c.out, nil
}

// OutputBounds returns a rectangle of the size of the intersection of the mask and source image. You can think of
// the rectangle dimensions as:
// 		width = min(src.width, mask.width) = rectangle.Max.X
//		height = min(src.height, mask.height) = rectangle.Max.Y
func (c *CCI) OutputBounds() image.Rectangle {
	topLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{}
	if c.src.Bounds().Max.X < c.mask.Bounds().Max.X {
		lowRight.X = c.src.Bounds().Max.X
	} else {
		lowRight.X = c.mask.Bounds().Max.X
	}
	if c.src.Bounds().Max.Y < c.mask.Bounds().Max.Y {
		lowRight.Y = c.src.Bounds().Max.Y
	} else {
		lowRight.Y = c.mask.Bounds().Max.Y
	}
	return image.Rectangle{Min: topLeft, Max: lowRight}
}

// getImagesOffset returns a pair of points representing the offset that must be applied when traversing the mask and
// source image within the intersection range. You may expect that maks and source images will be center-aligned.
// Points correspond to mask and source image offsets respectively.
func (c *CCI) getImagesOffset() (image.Point, image.Point) {
	fn := func(x1, x2 int) int {
		if x1 <= x2 {
			return 0
		}
		return (x1 - x2) / 2
	}
	mask := image.Point{
		X: fn(c.mask.Bounds().Max.X, c.src.Bounds().Max.X),
		Y: fn(c.mask.Bounds().Max.Y, c.src.Bounds().Max.Y),
	}
	src := image.Point{
		X: fn(c.src.Bounds().Max.X, c.mask.Bounds().Max.X),
		Y: fn(c.src.Bounds().Max.Y, c.mask.Bounds().Max.Y),
	}
	return mask, src
}

func isWhite(c color.Color) bool {
	r, g, b, a := c.RGBA()
	wR, wG, wB, wA := color.White.RGBA()
	return r == wR && g == wG && b == wB && a == wA
}
