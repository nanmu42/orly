[English](https://github.com/nanmu42/orly/blob/master/README.md) | **中文**

[![Build Status](https://github.com/nanmu42/orly/workflows/build/badge.svg)](https://github.com/nanmu42/orly/actions)
[![Lint Status](https://github.com/nanmu42/orly/workflows/golangci-lint/badge.svg)](https://github.com/nanmu42/orly/actions)
[![Docker Image](https://github.com/nanmu42/orly/workflows/Docker%20Image/badge.svg)](https://hub.docker.com/r/nanmu42/orly)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanmu42/orly)](https://goreportcard.com/report/github.com/nanmu42/orly)
[![GoDoc](https://godoc.org/github.com/nanmu42/orly?status.svg)](https://godoc.org/github.com/nanmu42/orly)

# O'RLY 封面工厂

O'RLY 封面工厂是一个娱乐性的书籍封面生成器，支持输入多种语言。 O'RLY 基于Golang和Vue.js实现。

你可以[点击这里](https://orly.nanmu.me/)创建你自己的O'RLY 封面，惊呆你的朋友和同事。

<table style="padding:10px">
  <tr>
    <td align="center"><img src="https://user-images.githubusercontent.com/8143068/47200455-d2d5e600-d3a8-11e8-8310-9125043b213a.jpeg" alt="ORLY: MySQL从入门到删数据库跑路" width=90%></td>
    <td align="center"><img src="https://user-images.githubusercontent.com/8143068/47200457-d2d5e600-d3a8-11e8-871b-2642a65565c4.jpeg" alt="ORLY: 脚本部署要义——不要多打空格" width=90%></td>
  </tr>
</table>

Docker镜像信息可以在[此处](https://hub.docker.com/r/nanmu42/orly)获取。

# Docker

```bash
docker pull registry.cn-hongkong.aliyuncs.com/nanmu42/orly:[tag]
```

可用tag请参阅[这里](https://github.com/nanmu42/orly/releases)。

示例配置：

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

保存为文件`rly.toml`，挂载到`/app/config`，启动`/app/rly -config config/rly.toml`.

# 开发

使用下列命令可编译 O'RLY：

```bash
mkdir -p assets
cd assets
wget -nc https://github.com/nanmu42/orly/releases/download/1.5.0-beta/cover-images.tar.xz
wget -nc https://github.com/nanmu42/orly/releases/download/1.1.0-beta/fonts.tar.xz
cd ..
make all
```

O'RLY 是一个简易的API服务和静态前端的组合：

* API服务的源代码位于 `cmd/rly`；
* 前端代码位于 `frontend`

动物图片以及字体文件分别位于`coverimage`和`font`之中。

您可以在这些文件夹中找到更详细的说明。

# 贡献

欢迎任何形式的贡献！

以下是一些推荐的方向：

* 为O'RLY的前端添加更多翻译；
* 改进O'RLY
* 为O'RLY增加新功能
* 提交Bug

您也可以

* 分享您在O'RLY上的作品和快乐 :wink:

# 贡献者列表

:hugs: 感谢以下贡献者：

* [TahsinGokalp](https://github.com/TahsinGokalp) （土耳其语翻译）
* [wooogi123](https://github.com/wooogi123) （韩语翻译）
* [cauldnz](https://github.com/cauldnz) （新增动物）

# 兄弟项目

与O'RLY功能类似而且很棒的项目：

* [dev.to上的O RLY Cover Generator](https://dev.to/rly) O'RLY 从这里汲取了大量的实现模式。这个服务还提供[slack集成](https://dev.to/rlyslack)。服务只支持英文，这是我决定开发O'RLY的部分原因。
* [日本的 O'Reilly Generator](http://oreilly-generator.com/) 封面在前端生成，用户体验很好。

# 注解

"O'RLY 封面工厂" 仅为娱乐，与O'Reilly Media没有任何联系。 

本作品使用了由Google和Adobe牵头制作的思源宋体以及思源黑体，
 中国的常州华文、韩国的 Sandoll Communications、日本的 Iwata Corporation
 参与了上述字体的设计制作。 
 
本作品使用了由Adobe牵头制作的Source Sans Pro，设计师为Paul D. Hunt.

封面所用图片来自于 [USF ClipArt ETC](https://etc.usf.edu/clipart/) 项目。
 
 # 致谢
 
 作者在此向[JetBrains](https://www.jetbrains.com/?from=ORLY-cover-generator)表示谢意，它提供了一份开源软件许可来支持作者的开源工作。
 
 # 项目许可
 
 本项目基于MIT协议开源。
 
 您可以在项目根目录找到一份协议内容拷贝。
