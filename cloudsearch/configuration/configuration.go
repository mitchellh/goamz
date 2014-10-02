package configuration

import (
	"net/http"
	"github.com/mitchellh/goamz/aws"
	"net/url"
	"encoding/xml"
	"strconv"
	awsauth "github.com/smartystreets/go-aws-auth"
)

type CloudsearchConfiguration struct {
	aws.Auth
	aws.Region
	httpClient *http.Client
}

const apiVersion = "2013-01-01"

func New(auth aws.Auth, region aws.Region) *CloudsearchConfiguration {
	return NewWithClient(auth, region, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, httpClient *http.Client) *CloudsearchConfiguration {
	return &CloudsearchConfiguration{auth, region, httpClient}
}

// Requests
type DomainStatus struct {
	ARN                    string
	Created                bool
	Deleted                bool
}
type CreateDomainResult struct {
	DomainStatus            string `xml:"DomainStatus"`
}

type ListDomainNamesResponse struct {
	XMLName                   xml.Name   `xml:"ListDomainNamesResponse"`
	ListDomainNamesResult     ListDomainNamesResult `xml:"ListDomainNamesResult"`
	RequestId                 string    `xml:"ResponseMetadata>RequestId"`
}

type ListDomainNamesResult struct {
	DomainNames              XmlMapStringString     `xml:"DomainNames"`
}

type XmlMapStringString struct {
	Entries                  []XmlMapStringStringEntry `xml:"entry"`
}

func (x XmlMapStringString) Map() map[string]string {
	res := make(map[string]string)
	for _, i := range x.Entries {
		res[i.Key] = i.Value
	}
	return res
}

type XmlMapStringStringEntry struct {
	Key                      string `xml:"key"`
	Value                    string `xml:"value"`
}

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
		prefix = e.Code+": "
	}
	if prefix == "" && e.StatusCode > 0 {
		prefix = strconv.Itoa(e.StatusCode)+": "
	}
	return prefix + e.Message
}

type xmlErrors struct {
	Errors []Error `xml:"Error"`
}

func (r *CloudsearchConfiguration) ListDomainNames() (*ListDomainNamesResponse, error) {

	params := map[string]string{}
	res := &ListDomainNamesResponse{}
	err := r.query("ListDomainNames", params, res)
	switch {
	case err != nil: return nil, err
	default: return res, nil
	}

}

func (r *CloudsearchConfiguration) CreateNewDomain(name string) (*CreateDomainResult, error) {
	params := map[string]string{"DomainName":name}
	err := r.query("CreateNewDomain", params, CreateDomainResult{})
	switch {
	case err != nil: return nil, err
	default: return nil, nil
	}
}


func (r *CloudsearchConfiguration) query(action string, params map[string]string, resp interface{}) error {
	endpoint, err := url.Parse(r.Region.CloudSearchEndpoint)

	query := endpoint.Query()
	query.Set("Version", apiVersion)
	query.Set("Action", action)
	for k, v := range params {
		query.Set(k, v)
	}
	endpoint.RawQuery = query.Encode()

	hReq, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		return err
	}

	awsauth.Sign(hReq, awsauth.Credentials{
		AccessKeyID: r.Auth.AccessKey,
		SecretAccessKey: r.Auth.SecretKey,
	})

	re, err := r.httpClient.Do(hReq)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	if re.StatusCode > 200 {
		return buildError(re)
	}
	decoder := xml.NewDecoder(re.Body)
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
