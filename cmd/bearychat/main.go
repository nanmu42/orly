/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by an MIT License.
 * You may find a license copy in project root.
 *
 */

// bearychat is a orly integration on Bearychat,
// a Slack-like working group.
//
// Most users speaks Chinese.
package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/nanmu42/orly/cmd/common"
	"github.com/pkg/errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/bearyinnovative/bearychat-go/openapi"

	"github.com/bearyinnovative/bearychat-go"
)

var (
	// config file location
	configFile *string
	logger     *zap.Logger
	// Version build params
	Version string
	// BuildDate build params
	BuildDate string
)

var (
	// colorSet corresponds with orly project
	colorSet = [...]string{
		"61005e",
		"70706d",
		"890029",
		"c4000e",
		"6d001d",
		"6a00bd",
		"f10000",
		"0071b1",
		"f9bc00",
		"2c0077",
		"ba009a",
		"009047",
		"009d9e",
		"222e85",
		"bd002e",
		"009d1a",
		"75a500",
	}
)

const (
	separator      = ';'
	alterSeparator = '；'
	srctag         = "bearychat"
	helpCommand    = "help"
	helpContent    = "用法：\n`@orly {标题};{顶部文字};{作者};[副标题];[图片序号 0-40];[颜色序号 0-16]`\n前三个参数必填，分隔符可以用中英文分号。标题可以用`Ctrl+回车`完成换行，其他参数中的回车会被忽略。\n私聊中可以忽略`@orly`\n图片序号和颜色需要可参考： https://rly.nanmu.me/"
	helloCommand   = "hello"
	helloContent   = "喵～ []~(￣▽￣)~*"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	configFile = flag.String("config", "config.toml", "config.toml file location for rly")
	w := common.NewBufferedLumberjack(&lumberjack.Logger{
		Filename:   "bearychat.log",
		MaxSize:    300, // megabytes
		MaxBackups: 5,
		MaxAge:     28, // days
	}, 32*1024)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(w),
		zap.InfoLevel,
	)
	logger = zap.New(core)
}

func main() {
	var err error
	defer logger.Sync()
	defer func() {
		if err != nil {
			fmt.Println(err)
			logger.Error("fatal error", zap.Error(err))
			os.Exit(1)
		}
	}()

	flag.Parse()
	fmt.Printf(`O'rly Generator API(%s)
built on %s

`, Version, BuildDate)

	err = C.LoadFrom(*configFile)
	if err != nil {
		err = errors.Wrap(err, "C.LoadFrom")
		return
	}

	botCtx, err := bearychat.NewRTMContext(C.RTMToken)
	if err != nil {
		err = errors.Wrap(err, "bearychat.NewRTMContext")
		return
	}

	err, messageChan, errChan := botCtx.Run()
	if err != nil {
		err = errors.Wrap(err, "botCtx.Run")
		return
	}

	botAPI := openapi.NewClient(C.RTMToken)

	var exitSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, exitSignals...)

	fmt.Println("Orly Bearychat Bot started :)")
	logger.Info("Orly Bearychat Bot started")

	for {
		select {
		case <-quitChan:
			fmt.Println("Orly Bearychat Bot is exiting safely...")
			logger.Info("Orly Bearychat Bot is exiting safely...")
			return
		case rtmErr := <-errChan:
			logger.Error("RTM error", zap.Error(rtmErr))
		case incoming := <-messageChan:
			if !incoming.IsChatMessage() || incoming.IsFromUID(botCtx.UID()) {
				continue
			}

			// only reply to mentioned
			if mentioned, content := incoming.ParseMentionUID(botCtx.UID()); mentioned {
				logger.Info("triggered",
					zap.Any("uid", incoming["uid"]),
					zap.Any("text", incoming["text"]),
				)

				vChannelID, ok := incoming["vchannel_id"].(string)
				if !ok {
					logger.Error("can not get vchannel_id")
					continue
				}

				content = strings.Trim(content, " \n\r\t")

				ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
				msgOpt := openapi.MessageCreateOptions{
					VChannelID: vChannelID,
				}
				switch true {
				case strings.ToLower(content) == helpCommand:
					msgOpt.Text = helpContent
				case strings.ToLower(content) == helloCommand:
					msgOpt.Text = helloContent
				default:
					imgURL, badImgRequest := conv(content)
					if badImgRequest != nil {
						logger.Info("bad img request", zap.Error(badImgRequest))
						msgOpt.Text = badImgRequest.Error() + "\n" + helpContent
					} else {
						logger.Info("generated", zap.String("url", imgURL))
						msgOpt.Text = "您的作品新鲜出炉 :D"
						msgOpt.Attachments = []openapi.MessageAttachment{
							{
								Images: []openapi.MessageAttachmentImage{
									{
										Url: &imgURL,
									},
								},
							},
						}
					}
				}
				outgoing, _, badLuck := botAPI.Message.Create(ctx, &msgOpt)
				cancel()
				if badLuck != nil {
					logger.Error("failed to make response", zap.Any("outgoing msg", outgoing), zap.Error(badLuck))
					continue
				}
			}
		}
	}
}

// conv parses request and construct image URL
func conv(content string) (URL string, err error) {
	content = consistentSeparator(content)
	parts := strings.Split(content, string(separator))
	paramLen := len(parts)
	if paramLen < 3 {
		err = errors.New("必填参数不足")
		return
	}
	values := url.Values{
		"src":      []string{srctag},
		"g_loc":    []string{"BR"},
		"title":    []string{parts[0]},
		"top_text": []string{removeLinefeed(parts[1])},
		"author":   []string{removeLinefeed(parts[2])},
	}
	// guide text
	if paramLen >= 4 {
		values.Set("g_text", removeLinefeed(parts[3]))
	} else {
		values.Set("g_text", "")
	}
	// image ID
	if paramLen >= 5 && strings.Trim(removeLinefeed(parts[4]), " ") != "" {
		imgID, badNum := strconv.ParseInt(removeLinefeed(parts[4]), 10, 64)
		if badNum != nil {
			err = errors.New("图片序号需要为数字哟")
			return
		}
		if imgID < 0 || imgID > 40 {
			err = errors.New("图片序号需要介于0~40之间")
			return
		}
		values.Set("img_id", strconv.FormatInt(imgID, 10))
	} else {
		values.Set("img_id", strconv.FormatInt(int64(rand.Intn(41)), 10))
	}
	// color ID
	if paramLen >= 6 && strings.Trim(removeLinefeed(parts[4]), " ") != "" {
		colorID, badNum := strconv.ParseInt(removeLinefeed(parts[4]), 10, 64)
		if badNum != nil {
			err = errors.New("颜色序号需要为数字哟")
			return
		}
		if colorID < 0 || colorID > 16 {
			err = errors.New("颜色序号需要介于0~16之间")
			return
		}
		values.Set("color", colorSet[colorID])
	} else {
		values.Set("color", colorSet[rand.Intn(17)])
	}

	URL = C.OrlyEndpoint + values.Encode()
	return
}

// consistentSeparator makes separator consistent
func consistentSeparator(content string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case alterSeparator:
			return ';'
		default:
			return r
		}
	}, content)
}

// removeLinefeed removes linefeed and space
func removeLinefeed(content string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '\n', '\r':
			return -1
		default:
			return r
		}
	}, content)
}
