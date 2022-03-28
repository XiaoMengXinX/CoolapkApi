package api

import (
	"fmt"
	coolapk "github.com/XiaoMengXinX/CoolapkApi-Go"
	"net/http"
	"strconv"
)

func FeedLike(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(getArg(r, "id"))

	api := coolapk.New()
	api.Cookie = r.Header.Get("Cookie")

	result, err := api.LikeFeed(id)
	if err != nil {
		w.WriteHeader(500)
	}

	for s, a := range result.Header {
		for _, i := range a {
			w.Header().Add(s, i)
		}
	}
	_, _ = fmt.Fprintf(w, result.Response)
}
