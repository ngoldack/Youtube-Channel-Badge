package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/internal"
)

func SubscriberCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	s := internal.ConvertToJson("Subscribe", internal.ChannelStats.SubscriberCount)
	fmt.Fprintf(w, s)
}
