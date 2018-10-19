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
CoverImageDir = "../../coverimage"
Debug = false
MaxImageID = 40
Port = ":3000"
TitleFont = "../../font/SourceHanSerif-Bold.ttc"
NormalFont = "../../font/SourceHanSans-Medium.ttc"
ORLYFont = "../../font/SourceHanSans-Heavy.ttc"
QueueLen = 20
WorkerNum = 4
Width = 800
```

* Have fun!
