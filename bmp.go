package gorimage

import (
	"github.com/disintegration/imaging"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type BMPImage struct {
	image  image.Image
	width  int
	height int
	source string
}

func NewBMPImage(file io.Reader, w, h int, s string, autoOrientation bool) (*BMPImage, error) {
	//img, err := bmp.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	imaging.Decode(file)
	if err != nil {
		return nil, err
	}

	return &BMPImage{width: w, height: h, source: s, image: img}, nil
}

func (i *BMPImage) Type() OutputImageType {
	return BMP
}

func (i *BMPImage) Size() (int, int) {
	return i.width, i.height
}

func (i *BMPImage) Data() image.Image {
	return i.image
}

func (i *BMPImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *BMPImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *BMPImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

func CreateBMPFile(dest string, data image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	//return bmp.Encode(out, data)
	return imaging.Encode(out, data, imaging.BMP)
}
