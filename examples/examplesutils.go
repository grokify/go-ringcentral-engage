package examples

import (
	"log"
	"net/http"

	"github.com/grokify/gotilla/fmt/fmtutil"
)

func HandleApiResponse(info interface{}, resp *http.Response, err error) {
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)
}
