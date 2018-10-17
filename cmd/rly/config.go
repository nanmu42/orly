/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	// C global settings
	C Setting
	// v viper
	v = viper.New()
)

// Setting is where config lies
type Setting struct {
	// cover width, height will be 1.4 * Width
	Width int
	// port string for gin like
	Port string
	// verbose mode
	Debug     bool
	WorkerNum int
	QueueLen  int
	// directory of cover image source file(*.tif)
	CoverImageDir string
	// for preheat
	MaxImageID int
	// path of fonts
	NormalFont string
	TitleFont  string
	ORLYFont   string
}

// AddPath adds path to config search scope
func (s *Setting) AddPath(path string) {
	v.AddConfigPath(path)
}

// LoadFrom loads config file from path
func (s *Setting) LoadFrom(path string) (err error) {
	v.SetConfigFile(path)
	err = v.ReadInConfig()
	if err != nil {
		err = errors.Wrap(err, "v.ReadInConfig")
		return
	}
	err = v.UnmarshalExact(s)
	if err != nil {
		err = errors.Wrap(err, "v.UnmarshalExact")
		return
	}
	info, err := s.Info()
	if err != nil {
		err = errors.Wrap(err, "s.Info")
		return
	}
	fmt.Printf("Using specified config file %s, content\n%s\n", path, info)
	return
}

// Load loads config file from default position path
func (s *Setting) Load() (err error) {
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("toml")
	err = v.ReadInConfig()
	if err != nil {
		err = errors.Wrap(err, "v.ReadInConfig")
		return
	}
	err = v.UnmarshalExact(s)
	if err != nil {
		err = errors.Wrap(err, "v.UnmarshalExact")
		return
	}
	info, err := s.Info()
	if err != nil {
		err = errors.Wrap(err, "s.Info")
		return
	}

	fmt.Printf("Using default config file, content:\n%s\n", info)

	return
}

// Info marshal setting into toml
func (s *Setting) Info() (info string, err error) {
	t, err := toml.Marshal(*s)
	if err != nil {
		err = errors.Wrap(err, "toml.Marshal")
		return
	}

	info = string(t)
	return
}
