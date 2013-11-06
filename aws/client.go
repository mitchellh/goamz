package aws

import (
	"log"
	"math"
	"net"
	"net/http"
	"time"
)

type Retriable func(*http.Request, *http.Response, error) bool

type Wait func(try int)

type ResilientTransport struct {
	// Timeout is the maximum amount of time a dial will wait for
	// a connect to complete.
	//
	// The default is no timeout.
	//
	// With or without a timeout, the operating system may impose
	// its own earlier timeout. For instance, TCP timeouts are
	// often around 3 minutes.
	DialTimeout time.Duration

	// ResponseHeaderTimeout, if non-zero, specifies the amount of
	// time to wait for a server's response headers after fully
	// writing the request (including its body, if any). This
	// time does not include the time to read the response body.
	ResponseHeaderTimeout time.Duration

	// RequestTimeout, if non-zero, specifies the amount of time for the entire
	// request. This includes dialing (if necessary), the response header as well
	// as the entire body.
	RequestTimeout time.Duration

	// MaxTries, if non-zero, specifies the number of times we will retry on
	// failure. Retries are only attempted for temporary network errors or known
	// safe failures.
	MaxTries    int
	Deadline    time.Time
	ShouldRetry Retriable
	Wait        Wait
	transport   *http.Transport
}

func NewClient(rt *ResilientTransport) *http.Client {
	rt.transport = &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			deadline := rt.Deadline
			c, err := net.DialTimeout(netw, addr, rt.DialTimeout)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(deadline)
			return c, nil
		},
		Proxy: http.ProxyFromEnvironment,
	}
	return &http.Client{
		Transport: rt,
	}
}

var retryingTransport = &ResilientTransport{
	Deadline:    time.Now().Add(5 * time.Second),
	DialTimeout: time.Second * 2,
	MaxTries:    3,
	ShouldRetry: awsRetry,
	Wait:        ExpBackoff,
}
var RetryingClient = NewClient(retryingTransport)

// TODO: I think I need to cancel requests that have timed out
func (t *ResilientTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	//return t.transport.RoundTrip(req)
	return t.tries(req)
}

func (t *ResilientTransport) tries(req *http.Request) (res *http.Response, err error) {
	for try := 0; try < t.MaxTries; try += 1 {
		log.Println("Try", try)
		res, err = t.transport.RoundTrip(req)

		if !t.ShouldRetry(req, res, err) {
			break
		}
		log.Println("Retrying ", try)

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
