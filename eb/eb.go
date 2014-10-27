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

// CreateConfigurationTemplate

type ConfigurationOptionSetting struct {
	Namespace  string
	OptionName string
	Value      string
}

type SourceConfiguration struct {
	ApplicationName string
	TemplateName    string
}

type CreateConfigurationTemplate struct {
	ApplicationName     string
	Description         string
	EnvironmentId       string
	OptionSettings      []ConfigurationOptionSetting
	SolutionStackName   string
	SourceConfiguration SourceConfiguration
	TemplateName        string
}

type CreateConfigurationTemplateResp struct {
	ApplicationName   string                       `xml:"CreateConfigurationTemplateResult>ApplicationName"`
	DateCreated       string                       `xml:"CreateConfigurationTemplateResult>DateCreated"`
	DateUpdated       string                       `xml:"CreateConfigurationTemplateResult>DateUpdated"`
	DeploymentStatus  string                       `xml:"CreateConfigurationTemplateResult>DeploymentStatus"`
	Description       string                       `xml:"CreateConfigurationTemplateResult>Description"`
	EnvironmentName   string                       `xml:"CreateConfigurationTemplateResult>EnvironmentName"`
	OptionSettings    []ConfigurationOptionSetting `xml:"CreateConfigurationTemplateResult>OptionSettings>member"`
	SolutionStackName string                       `xml:"CreateConfigurationTemplateResult>SolutionStackName"`
	TemplateName      string                       `xml:"CreateConfigurationTemplateResult>TemplateName"`
	RequestId         string                       `xml:"ResponseMetadata>RequestId"`
}

func (eb *EB) CreateConfigurationTemplate(options *CreateConfigurationTemplate) (resp *CreateConfigurationTemplateResp, err error) {
	params := makeParams("CreateConfigurationTemplate")

	params["ApplicationName"] = options.ApplicationName
	params["Description"] = options.Description
	params["EnvironmentId"] = options.EnvironmentId

	for i, v := range options.OptionSettings {
		params["OptionSettings.member."+strconv.Itoa(i+1)+"Namespace"] = v.Namespace
		params["OptionSettings.member."+strconv.Itoa(i+1)+"OptionName"] = v.OptionName
		params["OptionSettings.member."+strconv.Itoa(i+1)+"Value"] = v.Value
	}

	params["SolutionStackName"] = options.SolutionStackName
	params["SourceConfiguration.ApplicationName"] = options.SourceConfiguration.ApplicationName
	params["SourceConfiguration.TemplateName"] = options.SourceConfiguration.TemplateName
	params["TemplateName"] = options.TemplateName

	resp = &CreateConfigurationTemplateResp{}

	err = eb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// CreateEnvironment
// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_CreateEnvironment.html

type OptionSpecification struct {
	Namespace  string
	OptionName string
}

type Tag struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

type EnvironmentTier struct {
	Name    string
	Type    string
	Version string
}
type Listener struct {
	Port     int
	Protocol string
}
type LoadBalancerDescription struct {
	Domain           string
	Listeners        []Listener
	LoadBalancerName string
}

type EnvironmentResourcesDescription struct {
	LoadBalancer LoadBalancerDescription
}

type CreateEnvironment struct {
	ApplicationName   string
	CNAMEPrefix       string
	Description       string
	EnvironmentName   string
	OptionSettings    []ConfigurationOptionSetting
	OptionsToRemove   []OptionSpecification
	SolutionStackName string
	Tags              []Tag
	TemplateName      string
	Tier              EnvironmentTier
	VersionLabel      string
}

type CreateEnvironmentResp struct {
	ApplicationName   string                          `xml:"CreateEnvironmentResult>ApplicationName"`
	CNAME             string                          `xml:"CreateEnvironmentResult>CNAME"`
	DateCreated       string                          `xml:"CreateEnvironmentResult>DateCreated"`
	DateUpdated       string                          `xml:"CreateEnvironmentResult>DateUpdated"`
	Description       string                          `xml:"CreateEnvironmentResult>Description"`
	EndpointURL       string                          `xml:"CreateEnvironmentResult>EndpointURL"`
	EnvironmentId     string                          `xml:"CreateEnvironmentResult>EnvironmentId"`
	EnvironmentName   string                          `xml:"CreateEnvironmentResult>EnvironmentName"`
	Health            string                          `xml:"CreateEnvironmentResult>Health"`
	Resources         EnvironmentResourcesDescription `xml:"CreateEnvironmentResult>Resources"`
	SolutionStackName string                          `xml:"CreateEnvironmentResult>SolutionStackName"`
	Status            string                          `xml:"CreateEnvironmentResult>Status"`
	TemplateName      string                          `xml:"CreateEnvironmentResult>TemplateName"`
	Tier              EnvironmentTier                 `xml:"CreateEnvironmentResult>Tier"`
	VersionLabel      string                          `xml:"CreateEnvironmentResult>VersionLabel"`
	RequestId         string                          `xml:"ResponseMetadata>RequestId"`
}

func (eb *EB) CreateEnvironment(options *CreateEnvironment) (resp *CreateEnvironmentResp, err error) {
	params := makeParams("CreateEnvironment")

	params["ApplicationName"] = options.ApplicationName
	params["CNAMEPrefix"] = options.CNAMEPrefix
	params["Description"] = options.Description
	params["EnvironmentName"] = options.EnvironmentName

	for i, v := range options.OptionSettings {
		params["OptionSettings.member."+strconv.Itoa(i+1)+"Namespace"] = v.Namespace
		params["OptionSettings.member."+strconv.Itoa(i+1)+"OptionName"] = v.OptionName
		params["OptionSettings.member."+strconv.Itoa(i+1)+"Value"] = v.Value
	}

	for i, v := range options.OptionsToRemove {
		params["OptionsToRemove.member."+strconv.Itoa(i+1)+"Namespace"] = v.Namespace
		params["OptionsToRemove.member."+strconv.Itoa(i+1)+"OptionName"] = v.OptionName

	}
	params["SolutionStackName"] = options.SolutionStackName
	for i, v := range options.Tags {
		params["Tags.member."+strconv.Itoa(i+1)+"Key"] = v.Key
		params["Tags.member."+strconv.Itoa(i+1)+"Value"] = v.Value

	}
	params["TemplateName"] = options.TemplateName

	params["Tier.Name"] = options.Tier.Name
	params["Tier.Type"] = options.Tier.Type
	params["Tier.Version"] = options.Tier.Version

	params["VersionLabel"] = options.VersionLabel

	resp = &CreateEnvironmentResp{}

	err = eb.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// CreateStorageLocation

type CreateStorageLocationResp struct {
	S3Bucket  string `xml:"CreateStorageLocationResult>S3Bucket"`
	RequestId string `xml:"ResponseMetadata>RequestId"`
}

func (eb *EB) CreateStorageLocation() (resp *CreateStorageLocationResp, err error) {
	params := makeParams("CreateStorageLocation")

	resp = &CreateStorageLocationResp{}

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

// DeleteApplicationVersion
type DeleteApplicationVersion struct {
	ApplicationName    string
	DeleteSourceBundle bool
	VersionLabel       string
}

func (eb *EB) DeleteApplicationVersion(options *DeleteApplicationVersion) (resp *SimpleResp, err error) {
	params := makeParams("DeleteApplicationVersion")

	params["ApplicationName"] = options.ApplicationName
	params["VersionLabel"] = options.VersionLabel

	if options.DeleteSourceBundle {
		params["DeleteSourceBundle"] = "true"
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
