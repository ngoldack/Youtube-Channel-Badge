package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ChannelListResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID         string     `json:"id"`
	Statistics Statistics `json:"statistics"`
}

type Statistics struct {
	ViewCount       string `json:"viewCount"`
	CommentCount    string `json:"commentCount"`
	SubscriberCount string `json:"subscriberCount"`
	VideoCount      string `json:"videoCount"`
	Time            time.Time
}

var ChannelStats Statistics

func UpdateCounter() {
	reqString := fmt.Sprintf("https://www.googleapis.com/youtube/v3/channels?part=statistics&id=%s&key=%s", Info.ChannelID, Info.ApiKey)

	now := time.Now()
	if now.After(ChannelStats.Time.Add(time.Second * time.Duration(Info.CacheTime))) {
		fmt.Println("Not Fresh")
		resp, err := http.Get(reqString)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var r ChannelListResponse
		json.Unmarshal(body, &r)

		var channel Item
		for _, v := range r.Items {
			if v.ID == Info.ChannelID {
				channel = v
				break
			}
		}
		if channel.ID == "" {
			panic("no channel found with id: " + Info.ChannelID)
		}

		channel.Statistics.Time = time.Now()
		ChannelStats = channel.Statistics
	} else {
		fmt.Println("already fresh")
	}
}
