/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package orly

import (
	"log"
	"math/rand"
	"testing"

	"golang.org/x/image/colornames"

	"github.com/pkg/errors"
)

var coverFact *CoverFactory

const (
	maxImageID = 3
)

func init() {
	var err error

	provider := NewImageProvider(LoadTIFFFromFolder("coverimage"))
	normalFont, err := LoadFont("font/SourceHanSans-Medium.ttc")
	if err != nil {
		err = errors.Wrap(err, "LoadFont normalFont")
		return
	}
	titleFont, err := LoadFont("font/SourceHanSerif-Bold.ttc")
	if err != nil {
		err = errors.Wrap(err, "LoadFont titleFont")
		return
	}
	orlyFont, err := LoadFont("font/SourceHanSans-Heavy.ttc")
	if err != nil {
		err = errors.Wrap(err, "LoadFont orlyFont")
		return
	}

	coverFact = NewCoverFactory(1000, 1400, provider, titleFont, normalFont, orlyFont)
	// cache cover images
	err = coverFact.PreheatCache(maxImageID)
	if err != nil {
		err = errors.Wrap(err, "PreheatCache")
		log.Fatal(err)
	}
}

func BenchmarkCoverFactory_Draw(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		var err error
		for pb.Next() {
			_, err = coverFact.Draw("思源宋體",
				"Source Han Sans | 思源黑体 | 思源黑體 | 源ノ角ゴシック | 본고딕",
				"nanmu42",
				"思源黑體 | 源ノ角ゴシック | 본고딕",
				"BR",
				colornames.Red, rand.Intn(4))
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	if coverFact.CoverProvider.Miss() > maxImageID+1 {
		b.Fatalf("cache missed too much: got %v, want %v", coverFact.CoverProvider.Miss(), maxImageID+1)
	}
}
