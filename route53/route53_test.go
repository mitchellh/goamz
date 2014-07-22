package route53

import (
	"testing"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/testutil"
)

func makeTestServer() *testutil.HTTPServer {
	return testutil.NewHTTPServer()
}

func makeClient(server *testutil.HTTPServer) *Route53 {
	auth := aws.Auth{"abc", "123", ""}
	return NewWithClient(auth, aws.Region{Route53Endpoint: server.URL}, testutil.DefaultClient)
}

func TestCreateHostedZone(t *testing.T) {
	testServer := makeTestServer()
	testServer.Start()
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
}
