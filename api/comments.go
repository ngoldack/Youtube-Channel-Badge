package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/YoutubeChannelStats/internal"
)

func CommentCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	s := internal.ConvertToJson("Comments", internal.ChannelStats.CommentCount)
	fmt.Fprintf(w, s)
}
