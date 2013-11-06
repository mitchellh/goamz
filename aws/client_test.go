package aws_test

import (
	"fmt"
	"github.com/mwhooker/goamz/aws"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	//	"sync"
	"testing"
	"time"
)

func serveAndGet(handler http.HandlerFunc) (body string, err error) {
	ts := httptest.NewServer(handler)
	defer ts.Close()
	resp, err := aws.RetryingClient.Get(ts.URL)
	if err != nil {
		return
	}
	greeting, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	return strings.TrimSpace(string(greeting)), nil
}

func TestClient_expected(t *testing.T) {
	body := "foo bar"

	resp, err := serveAndGet(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, body)
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp != body {
		t.Fatal("Body not as expected.")
	}
}

func TestClient_delay(t *testing.T) {
	body := "baz"
	wait := 3
	resp, err := serveAndGet(func(w http.ResponseWriter, r *http.Request) {
		if wait < 0 {
			t.Fatal("Never succeeded.")
		}
		time.Sleep(time.Second * time.Duration(wait))
		wait -= 1
		fmt.Fprintln(w, body)
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp != body {
		t.Fatal("Body not as expected.", resp)
	}
}

func TestClient_retries(t *testing.T) {
	body := "biz"
	failed := false
	resp, err := serveAndGet(func(w http.ResponseWriter, r *http.Request) {
		if !failed {
			http.Error(w, "error", 500)
			failed = true
		} else {
			fmt.Fprintln(w, body)
		}
	})
	if failed != true {
		t.Error("We didn't retry!")
	}
	if err != nil {
		t.Fatal(err)
	}
	if resp != body {
		t.Fatal("Body not as expected.")
	}
}

func TestClient_fails(t *testing.T) {
	tries := 0
	_, err := serveAndGet(func(w http.ResponseWriter, r *http.Request) {
		tries += 1
		http.Error(w, "error", 500)
	})
	if err == nil {
		t.Fatal(err)
	}
	if tries != 3 {
		t.Fatal("Didn't retry enough")
	}
}
