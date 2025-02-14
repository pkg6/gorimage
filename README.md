# 特点

1.批量图像资源和转换
3.丰富的格式支持：JPG,JPEG,PNG,GIF,BMP,TIF,TIFF and WebP
4.多种重采样滤波器：NearestNeighbor, Box, Linear, Hermite, MitchellNetravali, CatmullRom, BSpline, Gaussian, Lanczos, Hann, Hamming, Blackman, Bartlett, Welch, Cosine

# 命令行

~~~
./gorimage --help
Usage of ./gorimage:
  -AutoOrientation
        If auto orientation is enabled, the image will be transformed after decoding according to the EXIF orientation tag (if present).
  -C string
        Process all image files in the directory
  -CPUMemUsage int
        Medium CPU & Memory cost, default process speed.
  -GIFNumColors int
        Maximum number of colors used in the GIF encoded image. NumColors ranges from 1 to 256 inclusive, higher is better.
  -JPEGQuality int
        Encoding parameter for JPG, JPEG images. Quality ranges from 1 to 100 inclusive, higher is better.
  -PNGCompression int
        Default Compression, No Compression, Best Speed, Best Compression
  -TIFFCompression int
        Uncompressed, Deflate, LZW, CCITTGroup3, CCITTGroup4
  -f string
        Processing individual image files
  -filter int
        Filter example:1,2,3,4,5,6,7,8,9,10,11,12,13,14,15
  -format int
        Format example: 1,2,3,4,5,6
  -height int
        Output image height
  -path string
        Output path will remain the same as the original file by default.
  -width int
        Output image width
~~~

