package wresize

import (
	"github.com/nfnt/resize"
	"image"
	"io"
)

func ThumbnailImage(img image.Image, width, height uint) image.Image {
	newImg := resize.Thumbnail(width, height, img, resize.Lanczos3)
	return newImg
}

// ThumbnailJpgFile
func ThumbnailFile(inputFile, outputFile string, width, height uint) error {
	img, err := OpenImageFile(inputFile)
	if err != nil {
		return err
	}
	newImg := ThumbnailImage(img, width, height)
	return WriteImageFile(outputFile, newImg)
}

// ThumbnailJpg
func OutThumbnailImage(in io.Reader, inFT FileType, width, height uint, outFT FileType, out io.Writer) error {
	inImg, err := OpenImage(in, inFT)
	if err != nil {
		return err
	}
	newImg := ThumbnailImage(inImg, width, height)
	return WriteImage(newImg, outFT, out)
}
