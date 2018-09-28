/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package orly

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"strconv"

	"github.com/pkg/errors"
)

const (
	// PaddingPctH padding in milli of width
	PaddingPctH = 40
	// ImageSizePctH image width and height in milli of height
	ImageSizePctH = 505
	// PrimaryBarHPct PrimaryBar Height in milli
	PrimaryBarHPct = 193
	// PrimaryBarPosPctH Y position of primary min point in milli
	PrimaryBarPosPctH = 573
	// SecondaryBarHPct secondaryBar Height in milli
	SecondaryBarHPct = 16

	// Denominator for ratio
	Denominator = 1e3
)

// Cover O RLY cover
type Cover struct {
	// TODO: add press name and press icon

	Width  int
	Height int

	// CoverProvider provides cached cover images
	CoverProvider *ImageProvider
}

// Draw outputs the cover in image
func (c *Cover) Draw( /*title string, topText string, author string, guideText string, guideTextPosition string, */ primaryColor color.Color, imageID int) (img *image.RGBA, err error) {
	img = image.NewRGBA(image.Rectangle{
		Min: image.ZP,
		Max: image.Point{c.Width, c.Height},
	})

	// fill cover with background color(white)
	draw.Draw(img, img.Rect, image.White, image.ZP, draw.Src)

	// draw two bars
	draw.Draw(img, c.secondaryBarRect(), image.NewUniform(primaryColor), image.ZP, draw.Src)
	draw.Draw(img, c.primaryBarRect(), image.NewUniform(primaryColor), image.ZP, draw.Src)

	// draw cover image
	coverRect := c.coverImgRect()
	coverSource, err := c.CoverProvider.Load(strconv.FormatInt(int64(imageID), 10)+".tif", coverRect)
	if err != nil {
		err = errors.Wrap(err, "c.CoverProvider.Load")
		return
	}
	fmt.Printf("need: w: %v px, h: %v px\ngot: w: %v px, h: %v px\n", coverRect.Dx(), coverRect.Dy(), coverSource.Bounds().Dx(), coverSource.Bounds().Dy())
	draw.Draw(img, coverRect, coverSource, c.coverPt(coverRect, coverSource.Bounds()), draw.Src)
	return
}

// coverPt calc proper cover src point
func (c *Cover) coverPt(dstRect image.Rectangle, imgRect image.Rectangle) (pt image.Point) {
	// a wide image should be middle in height
	if imgRect.Dx()*10/imgRect.Dy() > 15 {
		pt.X = imgRect.Dx() - dstRect.Dx()
		pt.Y = (imgRect.Dy() - dstRect.Dy()) / 2
	} else {
		pt.X = imgRect.Dx() - dstRect.Dx()
		pt.Y = imgRect.Dy() - dstRect.Dy()
	}
	return
}

// secondaryBarRect calc and return the rectangle of top bar
func (c *Cover) secondaryBarRect() image.Rectangle {
	var (
		coverPadding = PaddingPctH * c.Width / Denominator
		barHeight    = SecondaryBarHPct * c.Height / Denominator
		barWidth     = c.Width - (2 * coverPadding)
	)
	return image.Rectangle{
		Min: image.Point{
			X: coverPadding,
			Y: 0,
		},
		Max: image.Point{
			X: coverPadding + barWidth,
			Y: barHeight,
		},
	}

}

// secondaryBarRect calc and return the rectangle of title bar
func (c *Cover) primaryBarRect() image.Rectangle {
	var (
		coverPadding = PaddingPctH * c.Width / Denominator
		MinY         = PrimaryBarPosPctH * c.Height / Denominator
		barHeight    = PrimaryBarHPct * c.Height / Denominator
		barWidth     = c.Width - (2 * coverPadding)
	)
	return image.Rectangle{
		Min: image.Point{
			X: coverPadding,
			Y: MinY,
		},
		Max: image.Point{
			X: coverPadding + barWidth,
			Y: MinY + barHeight,
		},
	}
}

// coverImgRect calc and return the rectangle of cover image
func (c *Cover) coverImgRect() image.Rectangle {
	var (
		coverPadding = PaddingPctH * c.Width / Denominator
		MaxY         = PrimaryBarPosPctH * c.Height / Denominator
		imgHeight    = ImageSizePctH * c.Height / Denominator
	)
	return image.Rectangle{
		Min: image.Point{
			X: coverPadding,
			Y: MaxY - imgHeight,
		},
		Max: image.Point{
			X: c.Width - coverPadding,
			Y: MaxY,
		},
	}
}
