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
	"sync"
)

// ImageCache is an in-memory key-value image store
//
// ImageCache is safe for concurrency use.
type ImageCache struct {
	m sync.Map
}

// NewImageCache returns a empty image store
func NewImageCache() *ImageCache {
	return new(ImageCache)
}

// Store sets the image for a key.
func (i *ImageCache) Store(key string, img image.Image) {
	i.m.Store(key, img)
}

// Load returns the image stored in the map for a key,
// or nil if no value is present.
//
// The found result indicates whether value was found in the map.
func (i *ImageCache) Load(key string) (img image.Image, found bool) {
	thing, found := i.m.Load(key)
	if found {
		img = thing.(image.Image)
	}
	return
}
