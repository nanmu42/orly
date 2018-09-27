/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package orly

import (
	"image"
	"image/color"
	"image/draw"
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
	// Contents

	// Book title on big color bar
	Title string
	// text that goes up most
	TopText string
	// in down left corner
	Author string
	// around title
	GuideText string
	// predefined, four corner
	GuideTextPosition string

	// TODO: add picture, color, press name and press icon

	Width  int
	Height int

	// Color of bar
	PrimaryColor color.Color
}

// Draw outputs the cover in image
func (c *Cover) Draw() (img *image.RGBA, err error) {
	img = image.NewRGBA(image.Rectangle{
		Min: image.ZP,
		Max: image.Point{c.Width, c.Height},
	})

	// fill cover with background color(white)
	draw.Draw(img, img.Rect, image.White, image.ZP, draw.Src)

	// draw two bars
	draw.Draw(img, c.secondaryBarRect(), image.NewUniform(c.PrimaryColor), image.ZP, draw.Src)
	draw.Draw(img, c.primaryBarRect(), image.NewUniform(c.PrimaryColor), image.ZP, draw.Src)

	// draw cover image
	draw.Draw(img, c.coverImgRect(), image.NewUniform(color.RGBA{0xad, 0xff, 0x2f, 0xff}), image.ZP, draw.Src)
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
