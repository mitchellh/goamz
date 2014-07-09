// The autoscaling package provides types and functions for interaction with the AWS
// AutoScaling service (autoscaling)
package autoscaling

import (
	"encoding/xml"
	"github.com/mitchellh/goamz/aws"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// The AutoScaling type encapsulates operations operations with the autoscaling endpoint.
type AutoScaling struct {
	aws.Auth
	aws.Region
	httpClient *http.Client
}

const APIVersion = "2011-01-01"

// New creates a new AutoScaling instance.
func New(auth aws.Auth, region aws.Region) *AutoScaling {
	return NewWithClient(auth, region, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, httpClient *http.Client) *AutoScaling {
	return &AutoScaling{auth, region, httpClient}
}

func (autoscaling *AutoScaling) query(params map[string]string, resp interface{}) error {
	params["Version"] = APIVersion
	params["Timestamp"] = time.Now().In(time.UTC).Format(time.RFC3339)

	endpoint, err := url.Parse(autoscaling.Region.AutoScalingEndpoint)
	if err != nil {
		return err
	}

	sign(autoscaling.Auth, "GET", "/", params, endpoint.Host)
	endpoint.RawQuery = multimap(params).Encode()
	r, err := autoscaling.httpClient.Get(endpoint.String())

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
// AutoScaling objects

type Tag struct {
	Key               string `xml:"Key"`
	Value             string `xml:"Value"`
	PropogateAtLaunch string `xml:"PropogateAtLaunch"`
}

// ----------------------------------------------------------------------------
// Create

// The CreateAutoScalingGroup request parameters
type CreateAutoScalingGroup struct {
	AvailZone               []string
	DefaultCooldown         int
	DesiredCapacity         int
	HealthCheckGracePeriod  int
	HealthCheckType         string
	InstanceId              string
	LaunchConfigurationName string
	LoadBalancerNames       []string
	MaxSize                 int
	MinSize                 int
	PlacementGroup          string
	Name                    string
	Tags                    []Tag
	VPCZoneIdentifier       []string
}

func (autoscaling *AutoScaling) CreateAutoScalingGroup(options *CreateAutoScalingGroup) (resp *SimpleResp, err error) {
	params := makeParams("CreateAutoScalingGroup")

	params["AutoScalingGroupName"] = options.Name
	params["DefaultCooldown"] = strconv.Itoa(options.DefaultCooldown)
	params["DesiredCapacity"] = strconv.Itoa(options.DesiredCapacity)
	params["HealthCheckGracePeriod"] = strconv.Itoa(options.HealthCheckGracePeriod)
	params["HealthCheckType"] = options.HealthCheckType
	params["InstanceId"] = options.InstanceId
	params["LaunchConfigurationName"] = options.LaunchConfigurationName

	for i, v := range options.AvailZone {
		params["AvailabilityZones.member."+strconv.Itoa(i+1)] = v
	}

	for i, v := range options.LoadBalancerNames {
		params["LoadBalancerNames.member."+strconv.Itoa(i+1)] = v
	}

	params["MaxSize"] = strconv.Itoa(options.MaxSize)
	params["MinSize"] = strconv.Itoa(options.MinSize)

	params["PlacementGroup"] = options.PlacementGroup

	for j, tag := range options.Tags {
		params["Tag.member"+strconv.Itoa(j+1)+".Key"] = tag.Key
		params["Tag.member"+strconv.Itoa(j+1)+".Value"] = tag.Value
	}

	params["VPCZoneIdentifier"] = strings.Join(options.VPCZoneIdentifier, ",")

	resp = &SimpleResp{}

	err = autoscaling.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// The CreateLaunchConfiguration request parameters
type CreateLaunchConfiguration struct {
	ImageId        string
	InstanceId     string
	InstanceType   string
	KeyName        string
	Name           string
	SecurityGroups []string
}

func (autoscaling *AutoScaling) CreateLaunchConfiguration(options *CreateLaunchConfiguration) (resp *SimpleResp, err error) {
	params := makeParams("CreateLaunchConfiguration")

	params["LaunchConfigurationName"] = options.Name
	params["ImageId"] = options.ImageId
	params["InstanceType"] = options.InstanceType
	params["InstanceId"] = options.InstanceId

	for i, v := range options.SecurityGroups {
		params["SecurityGroups.member."+strconv.Itoa(i+1)] = v
	}

	resp = &SimpleResp{}

	err = autoscaling.query(params, resp)

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

// Error encapsulates an autoscaling error.
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
