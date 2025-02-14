package gorimage

import (
	"context"
	"fmt"
	"image"
	"log"
	"math"
	"path/filepath"
	"runtime"
	"strings"
)

var jobContext, cancel = context.WithCancel(context.Background())

func CancelBatchDealWith() {
	cancel()
	jobContext.Done()
}

func BatchDealWith(ctx context.Context, files []string, opts ImageOptions) {
	cpu := runtime.NumCPU()
	workerCount := math.Round(float64(cpu) * float64(opts.CPUMemUsage))
	swg := NewWaitGroup(int(workerCount))
	for _, f := range files {
		if err := swg.AddWithContext(jobContext); err != nil {
			fmt.Println(err)
			cancel()
		}
		go func() {
			defer swg.Done()
			if err := DealWithFile(ctx, f, opts); err != nil {
				fmt.Println(err)
			}
		}()
	}
	swg.Wait()
}

func DealWithFile(ctx context.Context, file string, option ImageOptions) error {
	logMsg := fmt.Sprintf("file: %s, options: %v", file, option)
	log.Println(logMsg)
	img, err := CreateImage(file, option.AutoOrientation)
	if err != nil {
		return err
	}
	dPath := img.Dir()
	if !strings.EqualFold(option.Path, "") {
		dPath = option.Path
	}
	dw, dh := img.Size()
	if option.Width > 0 {
		dw = option.Width
	}
	if option.Height > 0 {
		dh = option.Height
	}
	if option.Format == 0 {
		option.Format = img.Type()
	}
	img.Resize(dw, dh, option.Filter)
	_, err = writeFileByFormat(dPath, img.Name(), img.Data(), dw, dh, option)
	return err

}
func writeFileByFormat(path, name string, data image.Image,
	width, height int, option ImageOptions) (string, error) {
	file := fmt.Sprintf("%s_%dx%d", name, width, height)
	dest := filepath.Join(path, file)
	switch option.Format {
	case BMP:
		return dest + ExtBmp, CreateBMPFile(dest+ExtBmp, data)
	case GIF:
		return dest + ExtGif, CreateGIFFile(dest+ExtGif, data, option.GIFNumColors)
	case JPEG:
		return dest + ExtJpeg, CreateJPEGFile(dest+ExtJpeg, data, option.JPEGQuality)
	case JPG:
		return dest + ExtJpg, CreateJPEGFile(dest+ExtJpg, data, option.JPEGQuality)
	case PNG:
		return dest + ExtPng, CreatePNGFile(dest+ExtPng, data, option.PNGCompression)
	case TIFF:
		return dest + ExtTiff, CreateTIFFFile(dest+ExtTiff, data, option.TIFFCompression)
	case WEBP:
		//Webp will be transformed to PNGType, since I could not find a proper encoder for WebP
		return dest + ExtPng, CreatePNGFile(dest+ExtPng, data, option.PNGCompression)
	}
	return "", nil
}
