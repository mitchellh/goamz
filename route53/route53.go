// The elb package provides types and functions for interaction with the AWS
// Route53 service
package route53

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mitchellh/goamz/aws"
)

// The Route53 type encapsulates operations operations with the route53 endpoint.
type Route53 struct {
	aws.Auth
	aws.Region
	httpClient *http.Client
}

const APIVersion = "2013-04-01"

// New creates a new ELB instance.
func New(auth aws.Auth, region aws.Region) *Route53 {
	return NewWithClient(auth, region, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, httpClient *http.Client) *Route53 {
	return &Route53{auth, region, httpClient}
}

type CreateHostedZoneRequest struct {
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"HostedZoneConfig>Comment"`
}

type CreateHostedZoneResponse struct {
	HostedZone    HostedZone    `xml:"HostedZone"`
	ChangeInfo    ChangeInfo    `xml:"ChangeInfo"`
	DelegationSet DelegationSet `xml:"DelegationSet"`
}

type HostedZone struct {
	ID              string `xml:"Id"`
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"Config>Comment"`
	ResourceCount   int    `xml:"ResourceRecordSetCount"`
}

type ChangeInfo struct {
	ID          string `xml:"Id"`
	Status      string `xml:"Status"`
	SubmittedAt string `xml:"SubmittedAt"`
}

type DelegationSet struct {
	NameServers []string `xml:"NameServers>NameServer"`
}

func (r *Route53) query(method, path string, req, resp interface{}) error {
	params := make(map[string]string)
	endpoint, err := url.Parse(r.Region.Route53Endpoint)
	if err != nil {
		return err
	}
	endpoint.Path = path
	sign(r.Auth, endpoint.Path, params)

	// Encode the body
	var body io.ReadWriter
	if req != nil {
		body = bytes.NewBuffer(nil)
		enc := xml.NewEncoder(body)
		if err := enc.Encode(req); err != nil {
			return err
		}
	}

	// Make the http request
	hReq, err := http.NewRequest(method, endpoint.String(), body)
	if err != nil {
		return err
	}
	for k, v := range params {
		hReq.Header.Set(k, v)
	}
	re, err := r.httpClient.Do(hReq)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	// Check the status code
	switch re.StatusCode {
	case 200:
	case 201:
	default:
		return fmt.Errorf("Request failed, got status code: %d",
			re.StatusCode)
	}

	// Decode the response
	decoder := xml.NewDecoder(re.Body)
	return decoder.Decode(resp)
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}

// CreateHostedZone is used to create a new hosted zone
func (r *Route53) CreateHostedZone(req *CreateHostedZoneRequest) (*CreateHostedZoneResponse, error) {
	// Generate a unique caller reference if none provided
	if req.CallerReference == "" {
		req.CallerReference = time.Now().Format(time.RFC3339Nano)
	}
	out := &CreateHostedZoneResponse{}
	if err := r.query("POST", fmt.Sprintf("/%s/hostedzone", APIVersion), req, out); err != nil {
		return nil, err
	}
	return out, nil
}

type DeleteHostedZoneResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

func (r *Route53) DeleteHostedZone(ID string) (*DeleteHostedZoneResponse, error) {
	// Remove the hostedzone prefix if given
	ID = CleanZoneID(ID)
	out := &DeleteHostedZoneResponse{}
	err := r.query("DELETE", fmt.Sprintf("/%s/hostedzone/%s", APIVersion, ID), nil, out)
	if err != nil {
		return nil, err
	}
	return out, err
}

// CleanZoneID is used to remove the leading /hostedzone/
func CleanZoneID(ID string) string {
	if strings.HasPrefix(ID, "/hostedzone/") {
		ID = strings.TrimPrefix(ID, "/hostedzone/")
	}
	return ID
}

// CleanChangeID is used to remove the leading /change/
func CleanChangeID(ID string) string {
	if strings.HasPrefix(ID, "/change/") {
		ID = strings.TrimPrefix(ID, "/change/")
	}
	return ID
}

type GetHostedZoneResponse struct {
	HostedZone    HostedZone    `xml:"HostedZone"`
	DelegationSet DelegationSet `xml:"DelegationSet"`
}

func (r *Route53) GetHostedZone(ID string) (*GetHostedZoneResponse, error) {
	// Remove the hostedzone prefix if given
	ID = CleanZoneID(ID)
	out := &GetHostedZoneResponse{}
	err := r.query("GET", fmt.Sprintf("/%s/hostedzone/%s", APIVersion, ID), nil, out)
	if err != nil {
		return nil, err
	}
	return out, err
}

type GetChangeResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

func (r *Route53) GetChange(ID string) (string, error) {
	ID = CleanChangeID(ID)
	out := &GetChangeResponse{}
	err := r.query("GET", fmt.Sprintf("/%s/change/%s", APIVersion, ID), nil, out)
	if err != nil {
		return "", err
	}
	return out.ChangeInfo.Status, err
}
