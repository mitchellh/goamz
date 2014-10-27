// The eb package provides types and functions for interaction with the AWS
// Elastic Beanstalk service (EB)
package eb

import (
	"encoding/xml"
	"github.com/mitchellh/goamz/aws"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type EB struct {
	aws.Auth
	aws.Region
	httpClient *http.Client
}

const APIVersion = "2010-12-01"

// New creates a new ELB instance.
func New(auth aws.Auth, region aws.Region) *EB {
	return NewWithClient(auth, region, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, httpClient *http.Client) *EB {
	return &EB{auth, region, httpClient}
}

func (eb *EB) query(params map[string]string, resp interface{}) error {
	params["Version"] = APIVersion
	params["Timestamp"] = time.Now().In(time.UTC).Format(time.RFC3339)

	endpoint, err := url.Parse(eb.Region.EBEndpoint)
	if err != nil {
		return err
	}

	sign(eb.Auth, "GET", "/", params, endpoint.Host)
	endpoint.RawQuery = multimap(params).Encode()
	r, err := eb.httpClient.Get(endpoint.String())

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

// The CreateApplication request parameters
type CreateApplication struct {
	ApplicationName string
	Description     string
}

type CreateApplicationResp struct {
	ApplicationName        string   `xml:"CreateApplicationResult>Application>ApplicationName"`
	ConfigurationTemplates []string `xml:"CreateApplicationResult>Application>ConfigurationTemplates>member"`
	DateCreated            string   `xml:"CreateApplicationResult>Application>DateCreated"`
	DateUpdated            string   `xml:"CreateApplicationResult>Application>DateUpdated"`
	Description            string   `xml:"CreateApplicationResult>Application>Description"`
	Versions               []string `xml:"CreateApplicationResult>Application>Versions>member"`
	RequestId              string   `xml:"ResponseMetadata>RequestId"`
}

func (eb *EB) CreateApplication(options *CreateApplication) (resp *CreateApplicationResp, err error) {
	params := makeParams("CreateApplication")

	params["ApplicationName"] = options.ApplicationName
	params["Description"] = options.Description
	resp = &CreateApplicationResp{}

	err = eb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// ----------------------------------------------------------------------------
// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_CreateApplicationVersion.html

type S3Location struct {
	S3Bucket string
	S3Key    string
}

// The CreateApplicationVersion request parameters
type CreateApplicationVersion struct {
	ApplicationName       string
	AutoCreateApplication bool
	Description           string
	SourceBundle          S3Location
	VersionLabel          string
}

type CreateApplicationVersionResp struct {
	ApplicationName string     `xml:"CreateApplicationVersionResult>ApplicationVersion>ApplicationName"`
	DateCreated     string     `xml:"CreateApplicationVersionResult>ApplicationVersion>DateCreated"`
	DateUpdated     string     `xml:"CreateApplicationVersionResult>ApplicationVersion>DateUpdated"`
	Description     string     `xml:"CreateApplicationVersionResult>ApplicationVersion>Description"`
	SourceBundle    S3Location `xml:"CreateApplicationVersionResult>ApplicationVersion>SourceBundle"`
	VersionLabel    string     `xml:"CreateApplicationVersionResult>ApplicationVersion>VersionLabel"`
	RequestId       string     `xml:"ResponseMetadata>RequestId"`
}

func (eb *EB) CreateApplicationVersion(options *CreateApplicationVersion) (resp *CreateApplicationVersionResp, err error) {
	params := makeParams("CreateApplicationVersion")

	params["ApplicationName"] = options.ApplicationName
	params["Description"] = options.Description
	if options.AutoCreateApplication {
		params["AutoCreateApplication"] = "true"
	}
	params["SourceBundle.S3Bucket"] = options.SourceBundle.S3Bucket
	params["SourceBundle.S3Key"] = options.SourceBundle.S3Key
	params["VersionLabel"] = options.VersionLabel

	resp = &CreateApplicationVersionResp{}

	err = eb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// ----------------------------------------------------------------------------
// CheckDNSAvailability

type CheckDNSAvailability struct {
	CNAMEPrefix string
}

type CheckDNSAvailabilityResp struct {
	FullyQualifiedCNAME string `xml:"CheckDNSAvailabilityResult>FullyQualifiedCNAME"`
	Available           bool   `xml:"CheckDNSAvailabilityResult>Available"`
	RequestId           string `xml:"ResponseMetadata>RequestId"`
}

func (eb *EB) CheckDNSAvailability(options *CheckDNSAvailability) (resp *CheckDNSAvailabilityResp, err error) {
	params := makeParams("CheckDNSAvailability")

	params["CNAMEPrefix"] = options.CNAMEPrefix
	resp = &CheckDNSAvailabilityResp{}

	err = eb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// ----------------------------------------------------------------------------
// Delete

// The DeleteApplication request parameters
type DeleteApplication struct {
	ApplicationName     string
	TerminateEnvByForce bool
}

func (eb *EB) DeleteApplication(options *DeleteApplication) (resp *SimpleResp, err error) {
	params := makeParams("DeleteApplication")

	params["ApplicationName"] = options.ApplicationName

	if options.TerminateEnvByForce {
		params["TerminateEnvByForce"] = "true"
	}

	resp = &SimpleResp{}

	err = eb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

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
