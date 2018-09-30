/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"image/gif"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/nanmu42/orly"
	"golang.org/x/image/colornames"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	provider := orly.NewImageProvider(orly.LoadTIFFFromFolder("../../coverimage"))
	normalFont, err := orly.LoadFont("../../font/SourceHanSans-Medium.ttc")
	if err != nil {
		err = errors.Wrap(err, "LoadFont normalFont")
		return
	}
	titleFont, err := orly.LoadFont("../../font/SourceHanSerif-Bold.ttc")
	if err != nil {
		err = errors.Wrap(err, "LoadFont titleFont")
		return
	}
	orlyFont, err := orly.LoadFont("../../font/SourceHanSans-Heavy.ttc")
	if err != nil {
		err = errors.Wrap(err, "LoadFont orlyFont")
		return
	}

	cf := orly.NewCoverFactory(1000, 1400, provider, titleFont, normalFont, orlyFont)
	img, err := cf.Draw("思源宋體",
		"Source Han Sans | 思源黑体 | 思源黑體 | 源ノ角ゴシック | 본고딕",
		"nanmu42",
		"思源黑體 | 源ノ角ゴシック | 본고딕",
		"BR",
		colornames.Red, 0)
	if err != nil {
		return
	}

	f, err := os.Create("out.gif")
	if err != nil {
		return
	}
	defer f.Close()

	if err := gif.Encode(f, img, nil); err != nil {
		return
	}
}
