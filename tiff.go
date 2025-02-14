package gorimage

import (
	"github.com/disintegration/imaging"
	"golang.org/x/image/tiff"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type TIFFImage struct {
	image  image.Image
	source string
	width  int
	height int
}

func NewTIFFImage(file io.Reader, w, h int, s string, autoOrientation bool) (*TIFFImage, error) {
	//img, err := tiff.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	if err != nil {
		return nil, err
	}

	return &TIFFImage{width: w, height: h, source: s, image: img}, nil
}

func (i *TIFFImage) Size() (int, int) {
	return i.width, i.height
}

func (i *TIFFImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *TIFFImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *TIFFImage) Data() image.Image {
	return i.image
}

func (i *TIFFImage) Type() OutputImageType {
	return TIFF
}

func (i *TIFFImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

func CreateTIFFFile(dest string, data image.Image, compression int) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	return tiff.Encode(out, data, &tiff.Options{Compression: tiff.CompressionType(compression)})
}
