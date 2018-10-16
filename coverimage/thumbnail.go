// +build thumbnail

/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

// Thumbnail converts tiff image to png
package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io/ioutil"
	"os"
	"strings"

	"github.com/disintegration/imaging"

	"github.com/pkg/errors"
	"golang.org/x/image/tiff"
)

const (
	dirName   = "thumbnails"
	thumbSize = 120
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	err = os.RemoveAll(dirName)
	if err != nil {
		err = errors.Wrap(err, "os.RemoveAll")
		return
	}
	err = os.Mkdir(dirName, 0766)
	if err != nil {
		err = errors.Wrap(err, "os.Mkdir")
		return
	}

	info, err := ioutil.ReadDir(".")
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadDir")
		return
	}

	for _, item := range info {
		if !strings.HasSuffix(item.Name(), ".tif") {
			continue
		}
		err = thumbnail(item.Name())
		if err != nil {
			err = errors.Wrap(err, "thumbnail")
			return
		}
	}
}

func thumbnail(fileName string) (err error) {
	f, err := os.Open(fileName)
	if err != nil {
		err = errors.Wrap(err, "os.Open")
		return
	}
	defer f.Close()
	origin, err := tiff.Decode(f)
	if err != nil {
		err = errors.Wrapf(err, "tiff.Decode: %s", fileName)
		return
	}
	fitted := imaging.Fit(origin, thumbSize, thumbSize, imaging.Box)
	aligned := image.NewRGBA(image.Rect(0, 0, thumbSize, thumbSize))
	draw.Draw(aligned, aligned.Rect, image.White, image.ZP, draw.Src)
	draw.Draw(aligned, aligned.Rect, fitted, image.Pt(-(thumbSize-fitted.Bounds().Dx())/2, -(thumbSize-fitted.Bounds().Dy())/2), draw.Src)

	output, err := os.Create(dirName + string(os.PathSeparator) + fileName + ".gif")
	if err != nil {
		err = errors.Wrap(err, "os.Create")
		return
	}
	defer output.Close()

	err = gif.Encode(output, aligned, nil)
	if err != nil {
		err = errors.Wrap(err, "gif.Encode")
		return
	}
	return
}
