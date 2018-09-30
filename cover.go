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
	"strconv"
	"strings"

	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype"
	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"

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
	SecondaryBarHPct = 14

	// Denominator for ratio
	Denominator = 1e3

	// fonts

	// TitleSizePctH1 title font size of cover height in milli
	TitleSizePctH1 = 105
	// TitleSizePctH2 title(two lines) font size of cover height in milli
	TitleSizePctH2 = 55
	// TopTextSizePctH TopText font size of cover height in milli
	TopTextSizePctH = 20
	// AuthorSizePctH author font size of cover height in milli
	AuthorSizePctH = 24
	// GuideTextPctH GuideText font size of cover height in milli
	GuideTextPctH = 28
	// FontORLYPctH ORLY font size of cover height in milli
	FontORLYPctH = 36
)

// GuideText Position
const (
	BottomRight = "BR"
	BottomLeft  = "BL"
	TopRight    = "TR"
	TopLeft     = "TL"
)

// CoverFactory O RLY cover factory
type CoverFactory struct {
	width  int
	height int

	// CoverProvider provides cached cover images
	CoverProvider *ImageProvider
	// cover prototype
	CoverPrototype *image.RGBA

	// titleFont
	titleFont *truetype.Font
	// regularFont
	regularFont *truetype.Font
	// font for O RLY?
	orlyFont *truetype.Font
}

// NewCoverFactory initialize the cover
func NewCoverFactory(width, height int, provider *ImageProvider, titleFont, regularFont, orlyFont *truetype.Font) (c *CoverFactory) {
	prototype := image.NewRGBA(image.Rectangle{
		Min: image.ZP,
		Max: image.Point{width, height},
	})

	// fill cover with background color(white)
	draw.Draw(prototype, prototype.Rect, image.White, image.ZP, draw.Src)
	// draw O RLY?
	ctx := freetype.NewContext()

	ctx.SetClip(prototype.Rect)
	ctx.SetDst(prototype)

	ctx.SetHinting(font.HintingNone)
	ctx.SetSrc(image.Black)
	ctx.SetFont(orlyFont)
	ctx.SetFontSize(float64(FontORLYPctH * height / Denominator))
	var outPadding = PaddingPctH * height / Denominator
	ctx.DrawString("O RLY?", freetype.Pt(outPadding, height-outPadding))

	return &CoverFactory{
		width:          width,
		height:         height,
		CoverProvider:  provider,
		CoverPrototype: prototype,
		titleFont:      titleFont,
		regularFont:    regularFont,
		orlyFont:       orlyFont,
	}
}

// PreheatCache loads cover image into cache
func (c *CoverFactory) PreheatCache(maxImageID int) (err error) {
	coverRect := c.coverImgRect()
	for i := 0; i <= maxImageID; i++ {
		_, err = c.CoverProvider.Load(strconv.FormatInt(int64(i), 10)+".tif", coverRect)
		if err != nil {
			err = errors.Wrap(err, "CoverProvider.Load")
			return
		}
	}
	return
}

// Draw outputs the cover in image
func (c *CoverFactory) Draw(title, topText, author, guideText, guideTextPosition string, primaryColor color.Color, imageID int) (img *image.RGBA, err error) {
	img = image.NewRGBA(image.Rectangle{
		Min: image.ZP,
		Max: image.Point{c.width, c.height},
	})

	// fill cover with its prototype
	draw.Draw(img, img.Rect, c.CoverPrototype, image.ZP, draw.Src)

	// draw two bars
	theme := image.NewUniform(primaryColor)
	primaryBarRect := c.primaryBarRect()
	draw.Draw(img, primaryBarRect, theme, image.ZP, draw.Src)
	draw.Draw(img, c.secondaryBarRect(), theme, image.ZP, draw.Src)

	// draw cover image
	coverRect := c.coverImgRect()
	coverSource, err := c.CoverProvider.Load(strconv.FormatInt(int64(imageID), 10)+".tif", coverRect)
	if err != nil {
		err = errors.Wrap(err, "c.CoverProvider.Load")
		return
	}
	// fmt.Printf("need: w: %v px, h: %v px\ngot: w: %v px, h: %v px\n", coverRect.Dx(), coverRect.Dy(), coverSource.Bounds().Dx(), coverSource.Bounds().Dy())
	draw.Draw(img, coverRect, coverSource, c.coverPt(coverRect, coverSource.Bounds()), draw.Src)

	// prepare to draw letters
	ctx := freetype.NewContext()

	ctx.SetClip(img.Rect)
	ctx.SetDst(img)

	ctx.SetHinting(font.HintingNone)
	ctx.SetSrc(image.Black)
	ctx.SetFont(c.regularFont)

	var textWidth, lineHeight, textSize int
	var outPadding = PaddingPctH * c.height / Denominator
	// topText, mind the order
	textSize = TopTextSizePctH * c.height / Denominator
	lineHeight = textSize * 12 / 10
	ctx.SetFontSize(float64(textSize))
	textWidth = calcTextSize(ctx, topText)
	ctx.SetClip(img.Rect)
	ctx.SetDst(img)
	ctx.DrawString(topText, freetype.Pt((c.width-textWidth)/2, SecondaryBarHPct*c.height/Denominator+lineHeight))

	// GuideText, mind the order
	if len(guideText) > 0 {
		textSize = GuideTextPctH * c.height / Denominator
		lineHeight = textSize * 12 / 10
		ctx.SetFontSize(float64(textSize))
		textWidth = calcTextSize(ctx, guideText)
		ctx.SetClip(img.Rect)
		ctx.SetDst(img)
		var guideTextPivot fixed.Point26_6
		switch guideTextPosition {
		case BottomRight:
			guideTextPivot = freetype.Pt(primaryBarRect.Max.X-textWidth, primaryBarRect.Max.Y+lineHeight)
		case BottomLeft:
			guideTextPivot = freetype.Pt(primaryBarRect.Min.X, primaryBarRect.Max.Y+lineHeight)
		case TopRight:
			guideTextPivot = freetype.Pt(primaryBarRect.Max.X-textWidth, primaryBarRect.Min.Y-(lineHeight/6))
		case TopLeft:
			guideTextPivot = freetype.Pt(primaryBarRect.Min.X, primaryBarRect.Min.Y-(lineHeight/6))
		default:
			guideTextPivot = freetype.Pt(primaryBarRect.Max.X-textWidth, primaryBarRect.Max.Y+lineHeight)
		}
		ctx.DrawString(guideText, guideTextPivot)
	}

	// author, mind the order
	textSize = AuthorSizePctH * c.height / Denominator
	lineHeight = textSize * 12 / 10
	ctx.SetFontSize(float64(textSize))
	textWidth = calcTextSize(ctx, author)
	ctx.SetClip(img.Rect)
	ctx.SetDst(img)
	ctx.DrawString(author, freetype.Pt(c.width-textWidth-outPadding, c.height-outPadding))

	// title, line break is handled by user
	ctx.SetClip(img.Rect)
	ctx.SetDst(img)
	ctx.SetFont(c.titleFont)
	ctx.SetSrc(image.White)
	titleLines := strings.Split(title, "\n")
	switch len(titleLines) {
	case 1:
		textSize = TitleSizePctH1 * c.height / Denominator
		lineHeight = textSize * 12 / 10
		ctx.SetFontSize(float64(textSize))
		ctx.DrawString(titleLines[0], freetype.Pt(primaryBarRect.Min.X+outPadding/2, primaryBarRect.Max.Y-outPadding))
	case 2:
		textSize = TitleSizePctH2 * c.height / Denominator
		lineHeight = textSize * 12 / 10
		ctx.SetFontSize(float64(textSize))
		ctx.DrawString(titleLines[1], freetype.Pt(primaryBarRect.Min.X+outPadding/2, primaryBarRect.Max.Y-outPadding))
		ctx.DrawString(titleLines[0], freetype.Pt(primaryBarRect.Min.X+outPadding/2, primaryBarRect.Max.Y-outPadding-lineHeight))
	default:
		textSize = TitleSizePctH1 * c.height / Denominator
		lineHeight = textSize * 12 / 10
		ctx.SetFontSize(float64(textSize))
		ctx.DrawString(title, freetype.Pt(primaryBarRect.Min.X+outPadding/2, primaryBarRect.Max.Y-outPadding))
	}

	return
}

// coverPt calc proper cover src point
func (c *CoverFactory) coverPt(dstRect image.Rectangle, imgRect image.Rectangle) (pt image.Point) {
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
func (c *CoverFactory) secondaryBarRect() image.Rectangle {
	var (
		coverPadding = PaddingPctH * c.width / Denominator
		barHeight    = SecondaryBarHPct * c.height / Denominator
		barWidth     = c.width - (2 * coverPadding)
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
func (c *CoverFactory) primaryBarRect() image.Rectangle {
	var (
		coverPadding = PaddingPctH * c.width / Denominator
		MinY         = PrimaryBarPosPctH * c.height / Denominator
		barHeight    = PrimaryBarHPct * c.height / Denominator
		barWidth     = c.width - (2 * coverPadding)
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
func (c *CoverFactory) coverImgRect() image.Rectangle {
	var (
		coverPadding = PaddingPctH * c.width / Denominator
		MaxY         = PrimaryBarPosPctH * c.height / Denominator
		imgHeight    = ImageSizePctH * c.height / Denominator
	)
	return image.Rectangle{
		Min: image.Point{
			X: coverPadding,
			Y: MaxY - imgHeight,
		},
		Max: image.Point{
			X: c.width - coverPadding,
			Y: MaxY,
		},
	}
}

// calcTextSize returns estimated text width and height in pixel
// you need to reset ctx's dst and clip after this function call
func calcTextSize(ctx *freetype.Context, text string) int {
	temp := image.NewRGBA(image.Rect(0, 0, 0, 0))
	ctx.SetDst(temp)
	ctx.SetClip(image.White.Bounds())
	afterPt, err := ctx.DrawString(text, freetype.Pt(0, 0))
	if err != nil {
		return 0
	}
	return int(afterPt.X) >> 6
}
