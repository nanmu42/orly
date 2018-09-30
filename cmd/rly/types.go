/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"encoding/hex"
	"errors"
	"image/color"
	"net/url"
	"strconv"
)

// CoverQuery field name
const (
	titleField              = "title"
	topTextField            = "top_text"
	authorField             = "author"
	imageIDField            = "img_id"
	colorField              = "color"
	guideTextField          = "g_text"
	guideTextPlacementField = "g_loc"
)

// CoverQuery cover generation request
type CoverQuery struct {
	Title              string     `json:"title"`
	TopText            string     `json:"top_text"`
	Author             string     `json:"author"`
	ImageID            int64      `json:"image_id"`
	PrimaryColor       color.RGBA `json:"primary_color"`
	GuideText          string     `json:"guide_text"`
	GuideTextPlacement string     `json:"guide_text_placement"`
}

// ParseCoverQuery parse CoverQuery from request
//
// This function is written in an anti-DRY way to avoid using of reflect
func ParseCoverQuery(values url.Values) (cq CoverQuery, err error) {
	var (
		item []string
	)
	if item = values[titleField]; len(item) == 0 {
		err = lack(titleField)
		return
	}
	cq.Title = item[0]

	if item = values[topTextField]; len(item) == 0 {
		err = lack(topTextField)
		return
	}
	cq.TopText = item[0]

	if item = values[authorField]; len(item) == 0 {
		err = lack(authorField)
		return
	}
	cq.Author = item[0]

	if item = values[imageIDField]; len(item) == 0 {
		err = lack(imageIDField)
		return
	}
	cq.ImageID, err = strconv.ParseInt(item[0], 10, 64)
	if err != nil {
		err = errors.New("strconv.ParseInt: " + err.Error())
		return
	}

	if item = values[colorField]; len(item) == 0 {
		err = lack(colorField)
		return
	}
	// color is encoded as RRGGBB
	var colorBytes []byte
	colorBytes, err = hex.DecodeString(item[0])
	if err != nil || len(colorBytes) != 3 {
		err = errors.New("bad hex color")
		return
	}
	cq.PrimaryColor = color.RGBA{
		R: colorBytes[0],
		G: colorBytes[1],
		B: colorBytes[2],
		A: 0xff,
	}

	if item = values[guideTextField]; len(item) == 0 {
		err = lack(guideTextField)
		return
	}
	cq.GuideText = item[0]

	if item = values[guideTextPlacementField]; len(item) == 0 {
		err = lack(guideTextPlacementField)
		return
	}
	cq.GuideTextPlacement = item[0]

	return
}

// lack is a dummy error provider
func lack(field string) error {
	return errors.New("lack of field " + field)
}
