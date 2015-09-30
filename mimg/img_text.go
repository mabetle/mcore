package mimg

import (
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"strings"
)

func TextToImage(text string, fontSize float64) (image.Image, error) {

	fg, bg := image.White, image.Black
	rgba := image.NewRGBA(image.Rect(0, 0, 900, 900))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c := freetype.NewContext()
	//c.SetDPI(Dpi)
	//c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	// Draw the text.
	pt := freetype.Pt(10, 10)
	for _, s := range strings.Split(text, "\n") {
		_, err := c.DrawString(s, pt)
		if err != nil {
			return nil, err
		}
		//pt.Y += c.PointToFix32(12 * 1.5)
	}

	return rgba, nil
}
