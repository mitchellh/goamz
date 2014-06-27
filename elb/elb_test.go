package elb_test

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/elb"
	"github.com/mitchellh/goamz/testutil"
	. "github.com/motain/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
	elb *elb.ELB
}

var _ = Suite(&S{})

var testServer = testutil.NewHTTPServer()

func (s *S) SetUpSuite(c *C) {
	testServer.Start()
	auth := aws.Auth{"abc", "123", ""}
	s.elb = elb.NewWithClient(auth, aws.Region{ELBEndpoint: testServer.URL}, testutil.DefaultClient)
}

func (s *S) TearDownTest(c *C) {
	testServer.Flush()
}

func (s *S) TestCreateLoadBalancer(c *C) {
	testServer.Response(200, nil, CreateLoadBalancerExample)

	options := elb.CreateLoadBalancer{
		AvailZone: []string{"us-east-1a"},
		Listeners: []elb.Listener{elb.Listener{
			InstancePort:     80,
			InstanceProtocol: "http",
			LoadBalancerPort: 80,
			Protocol:         "http",
		},
		},
		LoadBalancerName: "foobar",
		Internal:         false,
		SecurityGroups:   []string{"sg1"},
		Subnets:          []string{"sn1"},
	}

	resp, err := s.elb.CreateLoadBalancer(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateLoadBalancer"})
	c.Assert(req.Form["LoadBalancerName"], DeepEquals, []string{"foobar"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "1549581b-12b7-11e3-895e-1334aEXAMPLE")
}

func (s *S) TestDeleteLoadBalancer(c *C) {
	testServer.Response(200, nil, DeleteLoadBalancerExample)

	options := elb.DeleteLoadBalancer{
		LoadBalancerName: "foobar",
	}

	resp, err := s.elb.DeleteLoadBalancer(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DeleteLoadBalancer"})
	c.Assert(req.Form["LoadBalancerName"], DeepEquals, []string{"foobar"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "1549581b-12b7-11e3-895e-1334aEXAMPLE")
}
