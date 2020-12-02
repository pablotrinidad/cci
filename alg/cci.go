package alg

import "image"

type CCI struct {
	src, mask image.Image
	seg       image.Image
}

func NewCCI(src, mask image.Image) *CCI {
	cci := &CCI{src: src, mask: mask}
	return cci
}

func (c *CCI) Run() {

}

func (c *CCI) SaveSegmentation(dest string) error {
	return nil
}
