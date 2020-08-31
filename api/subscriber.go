package api

import (
	"fmt"
	"net/http"

	"github.com/ntec.io/YoutubeChannelStats/internal"
)

func SubscriberCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	fmt.Fprintf(w, internal.ChannelStats.SubscriberCount)
}
