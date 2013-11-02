package aws

import (
	"github.com/mwhooker/go.httpcontrol"
	"net"
	"net/http"
	"time"
)

func awsRetry(req *http.Request, res *http.Response, err error) bool {
	retry := false

	if err == nil && res != nil {
		retry = false
	}
	if neterr, ok := err.(net.Error); ok {
		if neterr.Temporary() {
			retry = true
		}
	}
	if res != nil {
		if 500 > res.StatusCode && res.StatusCode >= 400 {
			retry = true
		}
	}
	return retry
}

var RetryingClient = &http.Client{
	Transport: &httpcontrol.Transport{
		Proxy:          http.ProxyFromEnvironment,
		RequestTimeout: time.Second * 5,
		DialTimeout:    time.Second * 2,
		ShouldRetry:    awsRetry,
		Wait:           httpcontrol.Wait(httpcontrol.ExpBackoff),
	},
}
