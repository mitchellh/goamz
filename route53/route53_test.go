package route53

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/testutil"
)

var testServer *testutil.HTTPServer

func init() {
	testServer = testutil.NewHTTPServer()
	testServer.Start()
}

func makeTestServer() *testutil.HTTPServer {
	testServer.Flush()
	log.Printf("Flush")
	return testServer
}

func makeClient(server *testutil.HTTPServer) *Route53 {
	auth := aws.Auth{"abc", "123", ""}
	return NewWithClient(auth, aws.Region{Route53Endpoint: server.URL}, testutil.DefaultClient)
}

func TestCreateHostedZone(t *testing.T) {
	testServer := makeTestServer()
	client := makeClient(testServer)
	testServer.Response(201, nil, CreateHostedZoneExample)

	req := &CreateHostedZoneRequest{
		Name:    "example.com",
		Comment: "Testing",
	}

	resp, err := client.CreateHostedZone(req)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if resp.HostedZone.ID != "/hostedzone/Z1PA6795UKMFR9" {
		t.Fatalf("bad: %v", resp)
	}
	if resp.ChangeInfo.ID != "/change/C1PA6795UKMFR9" {
		t.Fatalf("bad: %v", resp)
	}
	if resp.DelegationSet.NameServers[3] != "ns-2051.awsdns-67.co.uk" {
		t.Fatalf("bad: %v", resp)
	}

	httpReq := testServer.WaitRequest()
	if httpReq.URL.Path != "/2013-04-01/hostedzone" {
		t.Fatalf("bad: %#v", httpReq)
	}
	if httpReq.Method != "POST" {
		t.Fatalf("bad: %#v", httpReq)
	}
	if httpReq.ContentLength == 0 {
		t.Fatalf("bad: %#v", httpReq)
	}
}

func TestDeleteHostedZone(t *testing.T) {
	testServer := makeTestServer()
	client := makeClient(testServer)
	testServer.Response(200, nil, DeleteHostedZoneExample)

	resp, err := client.DeleteHostedZone("/hostedzone/foobarbaz")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if resp.ChangeInfo.ID != "/change/C1PA6795UKMFR9" {
		t.Fatalf("bad: %v", resp)
	}
}

func TestGetHostedZone(t *testing.T) {
	testServer := makeTestServer()
	client := makeClient(testServer)
	testServer.Response(200, nil, GetHostedZoneExample)

	resp, err := client.GetHostedZone("/hostedzone/foobarbaz")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if resp.HostedZone.CallerReference != "myUniqueIdentifier" {
		t.Fatalf("bad: %v", resp)
	}
}

func TestGetChange(t *testing.T) {
	testServer := makeTestServer()
	client := makeClient(testServer)
	testServer.Response(200, nil, GetChangeExample)

	status, err := client.GetChange("/change/abcd")
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if status != "INSYNC" {
		t.Fatalf("bad: %v", status)
	}
}

func decode(t *testing.T, r io.Reader, out interface{}) {
	var buf1 bytes.Buffer
	var buf2 bytes.Buffer
	b, err := io.Copy(io.MultiWriter(&buf1, &buf2), r)
	if err != nil {
		panic(fmt.Errorf("copy failed: %v", err))
	}
	if b == 0 {
		panic(fmt.Errorf("copy failed: zero bytes"))
	}
	dec := xml.NewDecoder(&buf1)
	if err := dec.Decode(out); err != nil {
		t.Errorf("body: %s||", buf2.Bytes())
		panic(fmt.Errorf("decode failed: %v", err))
	}
}
