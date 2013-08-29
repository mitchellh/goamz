package sqs_test

import (
	"crypto/md5"
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/exp/sqs"
	"github.com/mitchellh/goamz/testutil"
	"hash"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&S{})

type S struct {
	sqs *sqs.SQS
}

var testServer = testutil.NewHTTPServer()

func (s *S) SetUpSuite(c *C) {
	testServer.Start()
	auth := aws.Auth{"abc", "123", ""}
	s.sqs = sqs.New(auth, aws.Region{SQSEndpoint: testServer.URL})
}

func (s *S) TearDownTest(c *C) {
	testServer.Flush()
}

func (s *S) TestCreateQueue(c *C) {
	testServer.Response(200, nil, TestCreateQueueXmlOK)

	resp, err := s.sqs.CreateQueue("testQueue")
	req := testServer.WaitRequest()

	c.Assert(req.Method, Equals, "GET")
	c.Assert(req.URL.Path, Equals, "/")
	c.Assert(req.Header["Date"], Not(Equals), "")

	c.Assert(resp.Url, Equals, "http://sqs.us-east-1.amazonaws.com/123456789012/testQueue")
	c.Assert(err, IsNil)
}

func (s *S) TestListQueues(c *C) {
	testServer.Response(200, nil, TestListQueuesXmlOK)

	resp, err := s.sqs.ListQueues("")
	req := testServer.WaitRequest()

	c.Assert(req.Method, Equals, "GET")
	c.Assert(req.URL.Path, Equals, "/")
	c.Assert(req.Header["Date"], Not(Equals), "")

	c.Assert(len(resp.QueueUrl), Not(Equals), 0)
	c.Assert(resp.QueueUrl[0], Equals, "http://sqs.us-east-1.amazonaws.com/123456789012/testQueue")
	c.Assert(resp.ResponseMetadata.RequestId, Equals, "725275ae-0b9b-4762-b238-436d7c65a1ac")
	c.Assert(err, IsNil)
}

func (s *S) TestDeleteQueue(c *C) {
	testServer.Response(200, nil, TestDeleteQueueXmlOK)

	q := &sqs.Queue{s.sqs, testServer.URL + "/123456789012/testQueue/"}
	resp, err := q.Delete()
	req := testServer.WaitRequest()

	c.Assert(req.Method, Equals, "GET")
	c.Assert(req.URL.Path, Equals, "/123456789012/testQueue/")
	c.Assert(req.Header["Date"], Not(Equals), "")

	c.Assert(resp.ResponseMetadata.RequestId, Equals, "6fde8d1e-52cd-4581-8cd9-c512f4c64223")
	c.Assert(err, IsNil)
}

func (s *S) TestSendMessage(c *C) {
	testServer.Response(200, nil, TestSendMessageXmlOK)

	q := &sqs.Queue{s.sqs, testServer.URL + "/123456789012/testQueue/"}
	resp, err := q.SendMessage("This is a test message")
	req := testServer.WaitRequest()

	c.Assert(req.Method, Equals, "GET")
	c.Assert(req.URL.Path, Equals, "/123456789012/testQueue/")
	c.Assert(req.Header["Date"], Not(Equals), "")

	msg := "This is a test message"
	var h hash.Hash = md5.New()
	h.Write([]byte(msg))
	c.Assert(resp.MD5, Equals, fmt.Sprintf("%x", h.Sum(nil)))
	c.Assert(resp.Id, Equals, "5fea7756-0ea4-451a-a703-a558b933e274")
	c.Assert(err, IsNil)
}

func (s *S) TestReceiveMessage(c *C) {
	testServer.Response(200, nil, TestReceiveMessageXmlOK)

	q := &sqs.Queue{s.sqs, testServer.URL + "/123456789012/testQueue/"}
	resp, err := q.ReceiveMessage(5, 30)
	req := testServer.WaitRequest()

	c.Assert(req.Method, Equals, "GET")
	c.Assert(req.URL.Path, Equals, "/123456789012/testQueue/")
	c.Assert(req.Header["Date"], Not(Equals), "")

	c.Assert(len(resp.Messages), Not(Equals), 0)
	c.Assert(resp.Messages[0].MessageId, Equals, "5fea7756-0ea4-451a-a703-a558b933e274")
	c.Assert(resp.Messages[0].MD5OfBody, Equals, "fafb00f5732ab283681e124bf8747ed1")
	c.Assert(resp.Messages[0].ReceiptHandle, Equals, "MbZj6wDWli+JvwwJaBV+3dcjk2YW2vA3+STFFljTM8tJJg6HRG6PYSasuWXPJB+CwLj1FjgXUv1uSj1gUPAWV66FU/WeR4mq2OKpEGYWbnLmpRCJVAyeMjeU5ZBdtcQ+QEauMZc8ZRv37sIW2iJKq3M9MFx1YvV11A2x/KSbkJ0=")
	c.Assert(resp.Messages[0].Body, Equals, "This is a test message")
	c.Assert(len(resp.Messages[0].Attribute), Not(Equals), 0)
	c.Assert(err, IsNil)
}

func (s *S) TestChangeMessageVisibility(c *C) {
	testServer.Response(200, nil, TestReceiveMessageXmlOK)

	q := &sqs.Queue{s.sqs, testServer.URL + "/123456789012/testQueue/"}

	resp1, err := q.ReceiveMessage(1, 30)
	req := testServer.WaitRequest()

	testServer.Response(200, nil, TestChangeMessageVisibilityXmlOK)

	resp, err := q.ChangeMessageVisibility(&resp1.Messages[0], 50)
	req = testServer.WaitRequest()

	c.Assert(req.Method, Equals, "GET")
	c.Assert(req.URL.Path, Equals, "/123456789012/testQueue/")
	c.Assert(req.Header["Date"], Not(Equals), "")

	c.Assert(resp.ResponseMetadata.RequestId, Equals, "6a7a282a-d013-4a59-aba9-335b0fa48bed")
	c.Assert(err, IsNil)
}
