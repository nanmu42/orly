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
	"image/draw"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	"golang.org/x/image/tiff"
)

// ImageProvider is an loader and holder for assets of their various size.
// If the size asked does not exist, provider will load the origin file from
// file system, resize it and cache product in memory.
//
// ImageProvider is safe for concurrent use.
//
// For now, ImageProvider does not handle memory growing.
type ImageProvider struct {
	cache    *ImageCache
	loadLock *loadLock
	// originImgLoader if image does not exist,
	// provider calls this func to get original image for resize.
	originImgLoader func(fileName string) (image.Image, error)
	// miss record cache miss num
	miss uint64
}

// NewImageProvider factory of ImageProvider
func NewImageProvider(Loader func(fileName string) (image.Image, error)) *ImageProvider {
	return &ImageProvider{
		cache:           NewImageCache(),
		loadLock:        newLoadLock(),
		originImgLoader: Loader,
		miss:            0,
	}
}

// Load get image of specific size from cache or file system
func (i *ImageProvider) Load(fileName string, size image.Rectangle) (img image.Image, err error) {
	// check cache first
	key := encodeKey(fileName, size.Dx(), size.Dy())
	img, found := i.cache.Load(key)
	if found {
		return
	}
	// needs load from file system
	atomic.AddUint64(&i.miss, 1)
	// is loading?
	if i.loadLock.Lock(key) {
		err = errors.Errorf("%s is already under loading", key)
		return
	}
	// got the lock, now do loading
	defer i.loadLock.Unlock(key)
	origin, err := i.originImgLoader(fileName)
	if err != nil {
		err = errors.Wrap(err, "originImgLoader failed")
		return
	}
	// normally we are doing downscaling
	fitted := imaging.Fit(origin, size.Dx(), size.Dy(), imaging.Box)
	// converted to RGBA
	rgba := image.NewRGBA(fitted.Bounds())
	draw.Draw(rgba, rgba.Bounds(), fitted, image.Point{}, draw.Src)
	// push to cache
	i.cache.Store(key, rgba)
	img = rgba
	return
}

// Miss report the num of cache miss
func (i *ImageProvider) Miss() uint64 {
	return atomic.LoadUint64(&i.miss)
}

// LoadTIFFFromFolder func factory
func LoadTIFFFromFolder(folderPath string) func(fileName string) (image.Image, error) {
	return func(fileName string) (img image.Image, err error) {
		var strBuild strings.Builder
		strBuild.WriteString(folderPath)
		strBuild.WriteRune(os.PathSeparator)
		strBuild.WriteString(fileName)
		f, err := os.Open(strBuild.String())
		if err != nil {
			err = errors.Wrap(err, "Open failed")
			return
		}
		defer f.Close()
		img, err = tiff.Decode(f)
		if err != nil {
			err = errors.Wrap(err, "tiff.Decode")
			return
		}
		return
	}
}

// encodeKey encodes the key
func encodeKey(fileName string, width, height int) string {
	return strconv.FormatInt(int64(width), 10) + "&" + strconv.FormatInt(int64(height), 10) + "&" + fileName
}

// loadLock is a set of key-status pair which is concurrently safe.
type loadLock struct {
	mu   sync.Mutex
	dict map[string]*int64
}

// newLoadLock factory of loadLock
func newLoadLock() *loadLock {
	return &loadLock{
		mu:   sync.Mutex{},
		dict: make(map[string]*int64),
	}
}

// status get key's status pointer
func (m *loadLock) status(key string) *int64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.dict[key]; !ok {
		m.dict[key] = new(int64)
	}
	return m.dict[key]
}

// Lock locks key's loadlock
func (m *loadLock) Lock(key string) (alreadyLocked bool) {
	return !atomic.CompareAndSwapInt64(m.status(key), 0, 1)
}

// Unlock unlocks key's loadlock
func (m *loadLock) Unlock(key string) {
	atomic.StoreInt64(m.status(key), 0)
}
