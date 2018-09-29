package orly

import (
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
)

// LoadFont loads and parses font from file system
func LoadFont(fileName string) (loaded *truetype.Font, err error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		err = errors.Wrap(err, "ReadFile")
		return
	}
	loaded, err = freetype.ParseFont(fileBytes)
	if err != nil {
		err = errors.Wrap(err, "ParseFont")
		return
	}
	return
}
