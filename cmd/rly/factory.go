/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"github.com/nanmu42/orly"
	"github.com/pkg/errors"
)

// factory is the cover maker for requests
var factory *orly.CoverFactory

// initializeFactory makes factory available
func initializeFactory() (err error) {
	provider := orly.NewImageProvider(orly.LoadTIFFFromFolder(C.CoverImageDir))
	normalFont, err := orly.LoadFont(C.NormalFont)
	if err != nil {
		err = errors.Wrap(err, "LoadFont normalFont")
		return
	}
	titleFont, err := orly.LoadFont(C.TitleFont)
	if err != nil {
		err = errors.Wrap(err, "LoadFont titleFont")
		return
	}
	orlyFont, err := orly.LoadFont(C.ORLYFont)
	if err != nil {
		err = errors.Wrap(err, "LoadFont orlyFont")
		return
	}
	factory = orly.NewCoverFactory(C.Width, 14*C.Width/10, provider, titleFont, normalFont, orlyFont)
	err = factory.PreheatCache(C.MaxImageID)
	if err != nil {
		err = errors.Wrap(err, "PreheatCache")
		return
	}
	return
}

// short limit string to max length
func short(s string, max int) string {
	if len(s) > max {
		return s[:max]
	}
	return s
}
