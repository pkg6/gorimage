package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/pkg6/gorimage"
	"io/fs"
	"path/filepath"
)

type Flags struct {
	File   string
	Path   string
	Option gorimage.ImageOptions
	Format int
	Filter int
}

var flags Flags

func init() {
	flag.StringVar(&flags.File, "f", "", "Processing individual image files")
	flag.StringVar(&flags.Path, "C", "", "Process all image files in the directory")
	flag.IntVar(&flags.Format, "format", 0, "Format example: "+gorimage.OutputImagesTypesString())
	flag.IntVar(&flags.Filter, "filter", 0, "Filter example:"+gorimage.ResampleFilterTypesString())
	flag.StringVar(&flags.Option.Path, "path", "", "Output path will remain the same as the original file by default.")
	flag.IntVar(&flags.Option.Width, "width", 0, "Output image width")
	flag.IntVar(&flags.Option.Height, "height", 0, "Output image height")
	flag.IntVar(&flags.Option.JPEGQuality, "JPEGQuality", 0, "Encoding parameter for JPG, JPEG images. Quality ranges from 1 to 100 inclusive, higher is better.")
	flag.IntVar(&flags.Option.GIFNumColors, "GIFNumColors", 0, "Maximum number of colors used in the GIF encoded image. NumColors ranges from 1 to 256 inclusive, higher is better.")
	flag.IntVar(&flags.Option.TIFFCompression, "TIFFCompression", 0, "Uncompressed, Deflate, LZW, CCITTGroup3, CCITTGroup4")
	flag.IntVar(&flags.Option.PNGCompression, "PNGCompression", 0, "Default Compression, No Compression, Best Speed, Best Compression")
	flag.BoolVar(&flags.Option.AutoOrientation, "AutoOrientation", false, "If auto orientation is enabled, the image will be transformed after decoding according to the EXIF orientation tag (if present).")
	flag.IntVar(&flags.Option.CPUMemUsage, "CPUMemUsage", 0, "Medium CPU & Memory cost, default process speed.")
}
func main() {
	flag.Parse()
	if flags.Format > 0 {
		flags.Option.Format = gorimage.OutputImageType(flags.Format)
	}
	if flags.Filter > 0 {
		flags.Option.Filter = gorimage.ResampleFilterType(flags.Format)
	}
	var files []string
	if flags.File != "" {
		files = append(files, flags.File)
	}
	if flags.Path != "" {
		_ = filepath.Walk(flags.Path, func(path string, info fs.FileInfo, err error) error {
			if gorimage.IsSupportedImage(filepath.Ext(path)) {
				files = append(files, path)
			}
			return nil
		})
	}
	if len(files) <= 0 {
		fmt.Println("No files to process")
		return
	}
	gorimage.BatchDealWith(context.Background(), []string{flags.File}, flags.Option)
	defer gorimage.CancelBatchDealWith()
}
