package aws

import (
	"github.com/daaku/go.httpcontrol"
	"net/http"
	"time"
)

var RetryingClient = &http.Client{
	Transport: &httpcontrol.Transport{
		Proxy:          http.ProxyFromEnvironment,
		RequestTimeout: time.Second * 5,
		DialTimeout:    time.Second * 2,
		MaxTries:       3,
	},
}
