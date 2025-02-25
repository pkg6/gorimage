package gorimage

import (
	"github.com/disintegration/imaging"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type JPGImage struct {
	image  image.Image
	source string
	width  int
	height int
}

func NewJPGImage(file io.Reader, w, h int, s string, autoOrientation bool) (*JPGImage, error) {
	//img, err := jpeg.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	if err != nil {
		return nil, err
	}

	return &JPGImage{width: w, height: h, source: s, image: img}, nil
}

func (i *JPGImage) Size() (int, int) {
	return i.width, i.height
}

func (i *JPGImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *JPGImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *JPGImage) Data() image.Image {
	return i.image
}

func (i *JPGImage) Type() OutputImageType {
	return JPG
}

func (i *JPGImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

func CreateJPEGFile(dest string, data image.Image, quality int) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	//return jpeg.Encode(out, data, &jpeg.Options{Quality: quality})
	return imaging.Encode(out, data, imaging.JPEG, imaging.JPEGQuality(quality))
}
