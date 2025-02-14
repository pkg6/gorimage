package gorimage

import (
	"github.com/disintegration/imaging"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type GIFImage struct {
	image  image.Image
	source string

	width  int
	height int
}

func NewGIFImage(file io.Reader, w, h int, s string, autoOrientation bool) (*GIFImage, error) {
	//img, err := gif.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	if err != nil {
		return nil, err
	}

	return &GIFImage{width: w, height: h, source: s, image: img}, nil
}

func (i *GIFImage) Size() (int, int) {
	return i.width, i.height
}

func (i *GIFImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *GIFImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *GIFImage) Type() OutputImageType {
	return GIF
}

func (i *GIFImage) Data() image.Image {
	return i.image
}

func (i *GIFImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

func CreateGIFFile(dest string, data image.Image, colors int) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	//return gif.Encode(out, data, nil)
	return imaging.Encode(out, data, imaging.GIF, imaging.GIFNumColors(colors))
}
