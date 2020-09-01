package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/internal"
)

func VideoCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	s := internal.ConvertToJson("Videos", internal.ChannelStats.VideoCount)
	fmt.Fprintf(w, s)
}
