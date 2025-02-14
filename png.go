package gorimage

import (
	"github.com/disintegration/imaging"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type PNGImage struct {
	image  image.Image
	source string
	width  int
	height int
}

func NewPNGImage(file io.Reader, w, h int, s string, autoOrientation bool) (*PNGImage, error) {
	//img, err := png.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	if err != nil {
		return nil, err
	}

	return &PNGImage{width: w, height: h, source: s, image: img}, nil
}

func (i *PNGImage) Size() (int, int) {
	return i.width, i.height
}

func (i *PNGImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *PNGImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *PNGImage) Data() image.Image {
	return i.image
}

func (i *PNGImage) Type() OutputImageType {
	return PNG
}

func (i *PNGImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

func CreatePNGFile(dest string, data image.Image, compression int) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	//return png.Encode(out, data)
	return imaging.Encode(out, data, imaging.PNG, imaging.PNGCompressionLevel(png.CompressionLevel(compression)))
}
