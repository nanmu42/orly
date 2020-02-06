**English** | [中文](https://github.com/nanmu42/orly/blob/master/README_ZH.md)

[![Build Status](https://github.com/nanmu42/orly/workflows/build/badge.svg)](https://github.com/nanmu42/orly/actions)
[![Lint Status](https://github.com/nanmu42/orly/workflows/golangci-lint/badge.svg)](https://github.com/nanmu42/orly/actions)
[![Docker Image](https://github.com/nanmu42/orly/workflows/Docker%20Image/badge.svg)](https://hub.docker.com/r/nanmu42/orly)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanmu42/orly)](https://goreportcard.com/report/github.com/nanmu42/orly)
[![GoDoc](https://godoc.org/github.com/nanmu42/orly?status.svg)](https://godoc.org/github.com/nanmu42/orly)

# O'RLY Cover Generator

O'RLY Cover Generator is a parody book cover generator, 
implemented in Golang and Vue.js,
supporting a wide range of language including CJK.

O'RLY is living at https://orly.nanmu.me/

Docker images can be found at https://hub.docker.com/r/nanmu42/orly

# Docker Image

```bash
docker pull nanmu42/orly:[tag]
```

Refer to https://hub.docker.com/r/nanmu42/orly/tags for available tags.

# Develop O'RLY

O'RLY can be built in following commands:

```bash
mkdir -p assets
cd assets
wget -nc https://github.com/nanmu42/orly/releases/download/1.1.0-beta/cover-images.tar.xz
wget -nc https://github.com/nanmu42/orly/releases/download/1.1.0-beta/fonts.tar.xz
cd ..
make all
```

O'RLY consists of an API instance and a static frontend:

* API source lies in `cmd/rly`;
* Frontend source lies in `frontend`

Animal images and font files are in `coverimage` and `font`
respectively.

More documentations can be found in their directory.

# Contributes O'RLY

Contributions are always welcome!

Here are a few directions if you are interested:

* Help translating the frontend(we are using Vue i18n)
* Improve O'RLY
* Add New Features
* Raise a bug report

Or simply...

* Sharing the fun :wink:

# Projects in Brotherhood

There are several projects which share the idea:

* [O RLY Cover Generator on dev.to](https://dev.to/rly), where this project gets idea, supporting English only(partial reason for me to build O'RLY), there is also a [slack integration](https://dev.to/rlyslack). 
* [Japanese O'Reilly Generator](http://oreilly-generator.com/), with really good user experience, covers are generated in your browser

# Paperwork

"O'RLY Cover Generator" is just a parody, and it has no concern with O'Reilly Media. 

This work uses Source Han Serif and Source Han Sans from Adobe and Google, 
 with participation from partner foundries Changzhou SinoType in China, 
 Iwata Corporation in Japan, and Sandoll Communications in Korea. 
 
This work uses TrueType version of Source Sans Pro from Adobe by Paul D. Hunt.

The animal(well, not all of them are animal) images are from the [USF ClipArt ETC](https://etc.usf.edu/clipart/) project.

# License

Use of this work is governed by a MIT License.

You may find a license copy in project root.