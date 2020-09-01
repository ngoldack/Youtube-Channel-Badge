package api

import (
	"fmt"
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/internal"
)

func ViewCount(w http.ResponseWriter, req *http.Request) {
	internal.UpdateCounter()
	s := internal.ConvertToJson("Views", internal.ChannelStats.ViewCount)
	fmt.Fprintf(w, s)
}
