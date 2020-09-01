package main

import (
	"net/http"

	"github.com/ntec-io/Youtube-Channel-Badge/api"
)

func main() {
	http.HandleFunc("/subscribers", api.SubscriberCount)
	http.HandleFunc("/views", api.ViewCount)
	http.HandleFunc("/videos", api.VideoCount)
	http.HandleFunc("/comments", api.CommentCount)

	http.ListenAndServe(":8090", nil)
}
