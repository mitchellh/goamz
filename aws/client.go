package aws

import (
	"math"
	"net"
	"net/http"
	"time"
)

type Retriable func(*http.Request, *http.Response, error) bool

type Wait func(try int)

type ResilientTransport struct {
	// MaxTries, if non-zero, specifies the number of times we will retry on
	// failure. Retries are only attempted for temporary network errors or known
	// safe failures.
	MaxTries int

	ShouldRetry Retriable
	Wait        Wait
	transport   *http.Transport
}

var RetryingClient = &http.Client{
	Transport: &ResilientTransport{
		transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(5 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
			Proxy: http.ProxyFromEnvironment,
		},
		ShouldRetry: awsRetry,
		Wait:        Wait(ExpBackoff),
		MaxTries:    3,
	},
}

func (t *ResilientTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.tries(req)
}

func (t *ResilientTransport) tries(req *http.Request) (res *http.Response, err error) {
	for try := 0; try < t.MaxTries; try += 1 {
		res, err = t.transport.RoundTrip(req)

		if !t.ShouldRetry(req, res, err) {
			break
		}

		res.Body.Close()
		if t.Wait != nil {
			t.Wait(try)
		}
	}
	return
}

func ExpBackoff(try int) {
	time.Sleep(100 * time.Millisecond *
		time.Duration(math.Exp2(float64(try))))
}

func LinearBackoff(try int) {
	time.Sleep(100 * time.Millisecond * time.Duration(try))
}

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
		if 500 <= res.StatusCode && res.StatusCode < 600 {
			retry = true
		}
	}
	return retry
}
