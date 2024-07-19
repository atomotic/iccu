package nomi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/atomotic/iccu/client"
	"github.com/carlmjohnson/requests"
)

var base = "https://api.iccu.sbn.it/nomi/1.0.0/search?detail=full&nome=%s&page-size=10"

func Search(query string) *Response {
	ctx := context.Background()
	c := client.GetClient()

	q := url.QueryEscape(query)
	var r MainResponse
	err := requests.
		URL(fmt.Sprintf(base, q)).
		Client(c).
		ToJSON(&r).
		Fetch(ctx)

	if err != nil {
		fmt.Println(err)
	}
	return &r.Response

}
