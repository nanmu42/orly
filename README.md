**English** | [中文](https://github.com/nanmu42/orly/blob/master/README_ZH.md)

[![Build Status](https://github.com/nanmu42/orly/workflows/build/badge.svg)](https://github.com/nanmu42/orly/actions)
[![Lint Status](https://github.com/nanmu42/orly/workflows/golangci-lint/badge.svg)](https://github.com/nanmu42/orly/actions)
[![Docker Image](https://github.com/nanmu42/orly/workflows/Docker%20Image/badge.svg)](https://hub.docker.com/r/nanmu42/orly)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanmu42/orly)](https://goreportcard.com/report/github.com/nanmu42/orly)
[![GoDoc](https://godoc.org/github.com/nanmu42/orly?status.svg)](https://godoc.org/github.com/nanmu42/orly)

# O'RLY Cover Generator

O'RLY Cover Generator is a parody book cover generator, implemented in Golang and Vue.js, supporting a wide range of language including CJK.

[Go to the living demo](https://orly.nanmu.me/) to build your own O'RLY cover and troll your friends and colleagues.

<table style="padding:10px">
  <tr>
    <td align="center"><img src="https://user-images.githubusercontent.com/8143068/147319861-edaab37d-6e87-4bf3-84c4-034a7a8c9673.jpeg" alt="ORLY: Git Branch Rescue Guide" width=90%></td>
    <td align="center"><img src="https://user-images.githubusercontent.com/8143068/147319866-e6429227-a953-45b3-9984-ef85de0f4a94.jpeg" alt="ORLY: Float64 Math" width=90%></td>
  </tr>
</table>

Docker images can be found [here](https://hub.docker.com/r/nanmu42/orly).

# Docker Image

```bash
docker pull nanmu42/orly:[tag]
```

Refer to https://hub.docker.com/r/nanmu42/orly/tags for available tags.

Example config:

```toml
CoverImageDir = "cover-images"
Debug = false
MaxImageID = 41
Port = ":3000"
TitleFont = "fonts/SourceHanSerif-Bold.ttc"
NormalFont = "fonts/SourceHanSans-Medium.ttc"
ORLYFont = "fonts/SourceSansPro-Black.ttf"
QueueLen = 20
WorkerNum = 2
Width = 1000
```

Save as `rly.toml`, mount it into `/app/config`, and run docker image with param `/app/rly -config config/rly.toml`.

# Develop O'RLY

O'RLY can be built in following commands:

```bash
make assets
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

# Contributors

Many thanks :hugs: to following contributors:

* [TahsinGokalp](https://github.com/TahsinGokalp) (Turkish translation)
* [wooogi123](https://github.com/wooogi123) (Korean translation)
* [cauldnz](https://github.com/cauldnz) (Adding new animal)
* [Jessica Sachs](https://github.com/JessicaSachs) (Project maintainer!)

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

# Acknowledgements

The author would like to thank [JetBrains](https://www.jetbrains.com/?from=ORLY-cover-generator) for providing a JetBrains Open Source license for his open source developments.

# License

Use of this work is governed by the MIT License.

You may find a license copy in project root.
