//
// goamz - Go packages to interact with the Amazon Web Services.
//
//   https://wiki.ubuntu.com/goamz
//
// Copyright (c) 2011 Memeo Inc.
//
// Written by Prudhvi Krishna Surapaneni <me@prudhvi.net>
//
// This package is in an experimental state, and does not currently
// follow conventions and style of the rest of goamz or common
// Go conventions. It must be polished before it's considered a
// first-class package in goamz.
package sqs

import (
	"encoding/xml"
	"github.com/mitchellh/goamz/aws"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// The SQS type encapsulates operation with an SQS region.
type SQS struct {
	aws.Auth
	aws.Region
	private byte // Reserve the right of using private data.
}

func New(auth aws.Auth, region aws.Region) *SQS {
	return &SQS{auth, region, 0}
}

type Queue struct {
	*SQS
	Url string
}

type CreateQueueResponse struct {
	QueueUrl string `xml:"CreateQueueResult>QueueUrl"`
	ResponseMetadata
}

type ListQueuesResponse struct {
	QueueUrl []string `xml:"ListQueuesResult>QueueUrl"`
	ResponseMetadata
}

type DeleteMessageResponse struct {
	ResponseMetadata
}

type DeleteQueueResponse struct {
	ResponseMetadata
}

type SendMessageResponse struct {
	MD5 string `xml:"SendMessageResult>MD5OfMessageBody"`
	Id  string `xml:"SendMessageResult>MessageId"`
	ResponseMetadata
}

type ReceiveMessageResponse struct {
	Messages []Message `xml:"ReceiveMessageResult>Message"`
	ResponseMetadata
}

type Attribute struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type Message struct {
	MessageId     string      `xml:"MessageId"`
	Body          string      `xml:"Body"`
	MD5OfBody     string      `xml:"MD5OfBody"`
	ReceiptHandle string      `xml:"ReceiptHandle"`
	Attribute     []Attribute `xml:"Attribute"`
}

type ChangeMessageVisibilityResponse struct {
	ResponseMetadata
}

type GetQueueAttributesResponse struct {
	Attributes []Attribute `xml:"GetQueueAttributesResult>Attribute"`
	ResponseMetadata
}

type ResponseMetadata struct {
	RequestId string  `xml:"ResponseMetadata>RequestId"`
	BoxUsage  float64 `xml:"ResponseMetadata>BoxUsage"`
}

type Error struct {
	StatusCode int
	Code       string
	Message    string
	RequestId  string
}

func (err *Error) Error() string {
	return err.Message
}

type xmlErrors struct {
	RequestId string
	Errors    []Error `xml:"Errors>Error"`
}

func (s *SQS) CreateQueue(queueName string) (*Queue, error) {
	return s.CreateQueueWithTimeout(queueName, 30)
}

func (s *SQS) CreateQueueWithTimeout(queueName string, timeout int) (q *Queue, err error) {
	resp, err := s.newQueue(queueName, timeout)
	if err != nil {
		return nil, err
	}
	q = &Queue{s, resp.QueueUrl}
	return
}

func (s *SQS) QueueFromArn(queueUrl string) (q *Queue) {
	q = &Queue{s, queueUrl}
	return
}

func (s *SQS) newQueue(queueName string, timeout int) (resp *CreateQueueResponse, err error) {
	resp = &CreateQueueResponse{}
	params := makeParams("CreateQueue")

	params["QueueName"] = queueName
	params["DefaultVisibilityTimeout"] = strconv.Itoa(timeout)

	err = s.query("", params, resp)
	return
}

func (s *SQS) ListQueues(QueueNamePrefix string) (resp *ListQueuesResponse, err error) {
	resp = &ListQueuesResponse{}
	params := makeParams("ListQueues")

	if QueueNamePrefix != "" {
		params["QueueNamePrefix"] = QueueNamePrefix
	}

	err = s.query("", params, resp)
	return
}

func (q *Queue) Delete() (resp *DeleteQueueResponse, err error) {
	resp = &DeleteQueueResponse{}
	params := makeParams("DeleteQueue")

	err = q.SQS.query(q.Url, params, resp)
	return
}

func (q *Queue) SendMessage(MessageBody string) (resp *SendMessageResponse, err error) {
	resp = &SendMessageResponse{}
	params := makeParams("SendMessage")

	params["MessageBody"] = MessageBody

	err = q.SQS.query(q.Url, params, resp)
	return
}

func (q *Queue) ReceiveMessage(MaxNumberOfMessages, VisibilityTimeout int) (resp *ReceiveMessageResponse, err error) {
	resp = &ReceiveMessageResponse{}
	params := makeParams("ReceiveMessage")

	params["AttributeName"] = "All"
	params["MaxNumberOfMessages"] = strconv.Itoa(MaxNumberOfMessages)
	params["VisibilityTimeout"] = strconv.Itoa(VisibilityTimeout)

	err = q.SQS.query(q.Url, params, resp)
	return
}

func (q *Queue) ChangeMessageVisibility(M *Message, VisibilityTimeout int) (resp *ChangeMessageVisibilityResponse, err error) {
	resp = &ChangeMessageVisibilityResponse{}
	params := makeParams("ChangeMessageVisibility")
	params["VisibilityTimeout"] = strconv.Itoa(VisibilityTimeout)
	params["ReceiptHandle"] = M.ReceiptHandle

	err = q.SQS.query(q.Url, params, resp)
	return
}

func (q *Queue) GetQueueAttributes(A string) (resp *GetQueueAttributesResponse, err error) {
	resp = &GetQueueAttributesResponse{}
	params := makeParams("GetQueueAttributes")
	params["AttributeName"] = A

	err = q.SQS.query(q.Url, params, resp)
	return
}

func (q *Queue) DeleteMessage(M *Message) (resp *DeleteMessageResponse, err error) {
	resp = &DeleteMessageResponse{}
	params := makeParams("DeleteMessage")
	params["ReceiptHandle"] = M.ReceiptHandle

	err = q.SQS.query(q.Url, params, resp)
	return
}

func (s *SQS) query(queueUrl string, params map[string]string, resp interface{}) error {
	params["Timestamp"] = time.Now().UTC().Format(time.RFC3339)

	var err error
	var u *url.URL
	var path string

	if queueUrl != "" {
		u, err = url.Parse(queueUrl)
		path = queueUrl[len(s.Region.SQSEndpoint):]
	} else {
		u, err = url.Parse(s.Region.SQSEndpoint)
		path = "/"
	}

	if err != nil {
		return err
	}

	sign(s.Auth, "GET", path, params, u.Host)
	u.RawQuery = multimap(params).Encode()
	r, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return buildError(r)
	}

	err = xml.NewDecoder(r.Body).Decode(resp)
	return err
}

func buildError(r *http.Response) error {
	errors := xmlErrors{}
	xml.NewDecoder(r.Body).Decode(&errors)
	var err Error
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.RequestId = errors.RequestId
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

func makeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}
