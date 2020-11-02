# rly

`rly` is a cover generation API service.

# Build

```bash
./build.sh
```

# Run

* Init config file:

```bash
cp config_example.toml config.toml
```

* Edit `config.toml` to suit your need, e.g:

```toml
CoverImageDir = "cover-images"
Debug = false
MaxImageID = 40
Port = ":3000"
TitleFont = "fonts/SourceHanSerif-Bold.ttc"
NormalFont = "fonts/SourceHanSans-Medium.ttc"
ORLYFont = "fonts/SourceSansPro-Black.ttf"
QueueLen = 20
WorkerNum = 2
Width = 1000
```

Change font file path per your need.

* Have fun!
