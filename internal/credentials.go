package internal

import (
	"os"
	"strconv"
)

var Info Env

type Env struct {
	CacheTime int
	ApiKey    string
	ChannelID string
}

func init() {
	if s := os.Getenv("CACHE_TIME"); s != "" {
		Info.CacheTime, _ = strconv.Atoi(s)
		if Info.CacheTime <= 0 {
			Info.CacheTime = 300
		}
	} else {
		Info.CacheTime = 300
	}

	if Info.ApiKey = os.Getenv("API_KEY"); Info.ApiKey == "" {
		panic("api key empty")
	}
	if Info.ChannelID = os.Getenv("CHANNEL_ID"); Info.ChannelID == "" {
		panic("channel ID empty")
	}
}
