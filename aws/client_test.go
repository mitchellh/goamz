package aws_test

import (
	"fmt"
	"github.com/mwhooker/goamz/aws"
	"testing"
)

func TestOK(t *testing.T) {
	resp, err := aws.RetryingClient.Get("http://httpbin.org/")
	if err == nil {
		fmt.Println(resp, err)
	}
}

func TestDelay(t *testing.T) {
	resp, err := aws.RetryingClient.Get("http://httpbin.org/delay/6")
	if err == nil {
		fmt.Println(resp, err)
	}
}

func TestStatus(t *testing.T) {
	resp, err := aws.RetryingClient.Get("http://httpbin.org/status/500")
	if err == nil {
		fmt.Println(resp, err)
	}
}
