package internal

import (
	"os"
	"strconv"
)

var Info Env

type Env struct {
	SleepTime int
	ApiKey    string
	ChannelID string
}

func init() {
	if s := os.Getenv("SLEEP_TIME"); s != "" {
		Info.SleepTime, _ = strconv.Atoi(s)
		if Info.SleepTime <= 0 {
			Info.SleepTime = 300
		}
	} else {
		Info.SleepTime = 300
	}

	if Info.ApiKey = os.Getenv("API_KEY"); Info.ApiKey == "" {
		panic("api key empty")
	}
	if Info.ChannelID = os.Getenv("CHANNEL_ID"); Info.ChannelID == "" {
		panic("channel ID empty")
	}
}
