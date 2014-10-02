package configuration

import (
	"github.com/mitchellh/goamz/testutil"
	. "github.com/motain/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type CC struct {
	configuration *CloudsearchConfiguration
}

var _ = Suite(&CC{})

var testServer = testutil.NewHTTPServer()

func (s *CC) SetUpSuite(c *C) {
	testServer.Start()
	//	auth := aws.Auth{"abc", "123", ""}
	//	s.s3 = New(auth, aws.Region{Name: "faux-region-1", S3Endpoint: testServer.URL})
}



