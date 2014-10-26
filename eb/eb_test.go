package eb_test

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/eb"
	"github.com/mitchellh/goamz/testutil"
	. "github.com/motain/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
	eb *eb.EB
}

var _ = Suite(&S{})

var testServer = testutil.NewHTTPServer()

func (s *S) SetUpSuite(c *C) {
	testServer.Start()
	auth := aws.Auth{"abc", "123", ""}
	s.eb = eb.NewWithClient(auth, aws.Region{EBEndpoint: testServer.URL}, testutil.DefaultClient)
}

func (s *S) TearDownTest(c *C) {
	testServer.Flush()
}

func (s *S) TestCreateApplication(c *C) {
	testServer.Response(200, nil, CreateApplicationExample)

	options := eb.CreateApplication{
		ApplicationName: "SampleApp",
		Description:     "Sample Description",
	}

	resp, err := s.eb.CreateApplication(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateApplication"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["Description"], DeepEquals, []string{"Sample Description"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "8b00e053-f1d6-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDeleteLoadBalancer(c *C) {
	testServer.Response(200, nil, DeleteApplicationExample)

	options := eb.DeleteApplication{
		ApplicationName: "foobar",
	}

	resp, err := s.eb.DeleteApplication(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DeleteApplication"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"foobar"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "1f155abd-f1d7-11df-8a78-9f77047e0d0c")
}
