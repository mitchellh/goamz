// The elb package provides types and functions for interaction with the AWS
// Elastic Load Balancing service (ELB)
package elb

import (
	"encoding/xml"
	"github.com/mitchellh/goamz/aws"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// The ELB type encapsulates operations operations with the elb endpoint.
type ELB struct {
	aws.Auth
	aws.Region
	httpClient *http.Client
}

const APIVersion = "2012-06-01"

// New creates a new ELB instance.
func New(auth aws.Auth, region aws.Region) *ELB {
	return NewWithClient(auth, region, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, httpClient *http.Client) *ELB {
	return &ELB{auth, region, httpClient}
}

func (elb *ELB) query(params map[string]string, resp interface{}) error {
	params["Version"] = APIVersion
	params["Timestamp"] = time.Now().In(time.UTC).Format(time.RFC3339)

	endpoint, err := url.Parse(elb.ELBEndpoint)
	if err != nil {
		return err
	}

	sign(elb.Auth, "GET", "/", params, endpoint.Host)
	endpoint.RawQuery = multimap(params).Encode()
	r, err := elb.httpClient.Get(endpoint.String())

	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode > 200 {
		return buildError(r)
	}

	decoder := xml.NewDecoder(r.Body)
	decodedBody := decoder.Decode(resp)

	return decodedBody
}

func buildError(r *http.Response) error {
	var (
		err    Error
		errors xmlErrors
	)
	xml.NewDecoder(r.Body).Decode(&errors)
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}

func makeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}

// ----------------------------------------------------------------------------
// Create

// A listener attaches to an elb
type Listener struct {
	InstancePort     int64  `xml:"instancePort"`
	InstanceProtocol string `xml:"instanceProtocol"`
	LoadBalancerPort int64  `xml:"loadBalancerPort"`
	Protocol         string `xml:"protocol"`
}

// The CreateLoadBalancer request parameters
type CreateLoadBalancer struct {
	AvailZone        []string
	Listeners        []Listener
	LoadBalancerName string
	Internal         bool // true for vpc elbs
	SecurityGroups   []string
	Subnets          []string
}

type CreateLoadBalancerResp struct {
	DNSName   string `xml:"CreateLoadBalancerResult>DNSName"`
	RequestId string `xml:"ResponseMetadata>RequestId"`
}

func (elb *ELB) CreateLoadBalancer(options *CreateLoadBalancer) (resp *CreateLoadBalancerResp, err error) {
	params := makeParams("CreateLoadBalancer")

	params["LoadBalancerName"] = options.LoadBalancerName

	for i, v := range options.SecurityGroups {
		params["AvailabilityZones.member."+strconv.Itoa(i)] = v
	}

	for i, v := range options.AvailZone {
		params["SecurityGroups.member."+strconv.Itoa(i)] = v
	}

	for i, v := range options.Subnets {
		params["Subnets.member."+strconv.Itoa(i)] = v
	}

	for i, v := range options.Listeners {
		params["Subnets.member."+strconv.Itoa(i)+".LoadBalancerPort"] = strconv.FormatInt(v.LoadBalancerPort, 10)
		params["Subnets.member."+strconv.Itoa(i)+".InstancePort"] = strconv.FormatInt(v.InstancePort, 10)
		params["Subnets.member."+strconv.Itoa(i)+".Protocol"] = v.Protocol
		params["Subnets.member."+strconv.Itoa(i)+".InstanceProtocol"] = v.InstanceProtocol
	}

	if options.Internal {
		params["Scheme"] = "internal"
	}

	resp = &CreateLoadBalancerResp{}

	err = elb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// ----------------------------------------------------------------------------
// Destroy

// The DestroyLoadBalancer request parameters
type DeleteLoadBalancer struct {
	LoadBalancerName string
}

func (elb *ELB) DeleteLoadBalancer(options *DeleteLoadBalancer) (resp *SimpleResp, err error) {
	params := makeParams("DeleteLoadBalancer")

	params["LoadBalancerName"] = options.LoadBalancerName

	resp = &SimpleResp{}

	err = elb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// ----------------------------------------------------------------------------
// Responses

type SimpleResp struct {
	RequestId string `xml:"ResponseMetadata>RequestId"`
}

type xmlErrors struct {
	Errors []Error `xml:"Error"`
}

// Error encapsulates an elb error.
type Error struct {
	// HTTP status code of the error.
	StatusCode int

	// AWS code of the error.
	Code string

	// Message explaining the error.
	Message string
}

func (e *Error) Error() string {
	var prefix string
	if e.Code != "" {
		prefix = e.Code + ": "
	}
	if prefix == "" && e.StatusCode > 0 {
		prefix = strconv.Itoa(e.StatusCode) + ": "
	}
	return prefix + e.Message
}
