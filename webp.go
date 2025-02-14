package gorimage

import (
	"github.com/disintegration/imaging"
	"golang.org/x/image/webp"
	"image"
	"io"
	"path/filepath"
	"strings"
)

type WebPImage struct {
	image  image.Image
	source string
	width  int
	height int
}

func NewWebPImage(file io.Reader, w, h int, s string) (*WebPImage, error) {
	img, err := webp.Decode(file)
	if err != nil {
		return nil, err
	}

	return &WebPImage{width: w, height: h, source: s, image: img}, nil
}

func (i *WebPImage) Size() (int, int) {
	return i.width, i.height
}

func (i *WebPImage) Type() OutputImageType {
	return WEBP
}

func (i *WebPImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *WebPImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}
func (i *WebPImage) Data() image.Image {
	return i.image
}

func (i *WebPImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}
