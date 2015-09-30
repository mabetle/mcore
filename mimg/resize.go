package mimg

import (
	"github.com/nfnt/resize"
	"image"
	"io"
)

func ResizeImage(img image.Image, width, height uint) image.Image {
	newImg := resize.Resize(width, height, img, resize.Lanczos3)
	return newImg
}

// ResizeJpgFile
func ResizeFile(inputFile, outputFile string, width, height uint) error {
	img, err := OpenImageFile(inputFile)
	if err != nil {
		return err
	}
	newImg := ResizeImage(img, width, height)
	return WriteImageFile(outputFile, newImg)
}

// ResizeJpg
func OutResizeImage(in io.Reader, inFT FileType, width, height uint, outFT FileType, out io.Writer) error {
	inImg, err := OpenImage(in, inFT)
	if err != nil {
		return err
	}
	newImg := ResizeImage(inImg, width, height)
	return WriteImage(newImg, outFT, out)
}
